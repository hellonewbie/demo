package four

import "sync"

func search2(nums []int, target int) int {
	var wg sync.WaitGroup
	wg.Add(2)
	lef := 0
	rig := 0

	go func() {
		defer wg.Done()
		for i := 0; i < len(nums); i++ {
			if nums[i] == target {
				lef = i
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] == target {
				rig = i
				break
			}
		}
	}()
	wg.Wait()
	if rig+lef != -2 {
		return rig - lef + 1
	}
	return 0

}
