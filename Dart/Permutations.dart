class Solution
{
    List<List<int>> allPermutations = List.empty(growable: true);
    void swap(List<int> permutations, int l, int r)
    {
        int tmp = permutations[l];
        permutations[l] = permutations[r];
        permutations[r] = tmp;
    }
    
    void getAllPermutation(List<int> nums, int n)
    {
        if (n == nums.length - 1)
        {
            allPermutations.add([...nums]);
            return;
        }
        for (int i = n; i < nums.length; ++i)
        {
            swap(nums, i, n);
            getAllPermutation(nums, n + 1);
            swap(nums, n, i);
        }
    }
    List<List<int>> permute(List<int> nums)
    {
        getAllPermutation(nums, 0);
        return allPermutations;
    }
}
