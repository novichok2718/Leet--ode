class Solution
{
	bool isSortedDescending(List<int> nums)
	{
		for (int i = 0; i < nums.length - 1; ++i)
		{
			if (nums[i] < nums[i + 1]) return false;
		}
		return true;
	}

	void swap(List<int> nums, int target)
	{
		int min = 101;
		int index = 0;
		int tmp = 0;
		for (int i = 0; i < nums.length; ++i)
		{
			tmp = nums[i] - target;
			if (min > tmp && tmp > 0)
			{
				min = tmp;
				index = i;
			}
		}
		tmp = nums[0];
		nums[0] = nums[index];
		nums[index] = tmp;
	}

	void nextPermutation(List<int> nums)
	{
		List<int> window = [];
		for (var digit in nums.reversed)
		{
			window.insert(0, digit);
			if (!isSortedDescending(window))
			{
				swap(window, window[0]);	
				var tmp = window.removeAt(0);
				window.sort();
				window.insert(0, tmp);
				for (int i = nums.length - window.length; i < nums.length; ++i)
				{
					nums[i] = window[i - nums.length + window.length];
				}
				return;
			}
		}
		nums.sort();
	}
}
