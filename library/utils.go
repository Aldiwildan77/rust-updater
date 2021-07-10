package library

import "sync"

func Merge(errChans ...<-chan error) <-chan error {
	mergedChan := make(chan error)

	var wg sync.WaitGroup
	wg.Add(len(errChans))
	go func() {
		wg.Wait()
		close(mergedChan)
	}()

	for i := range errChans {
		go func(errChan <-chan error) {
			for err := range errChan {
				if err != nil {
					mergedChan <- err
				}
			}
			wg.Done()
		}(errChans[i])
	}

	return mergedChan
}
