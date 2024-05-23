class Solution 
{
    List<int> findSubstring(String s, List<String> words) 
    {
        List<int> startOfSubString = List<int>.empty(growable: true);
        Map<String, int> obj = {};
        for (var word in words)
        {
            int count = RegExp(RegExp.escape(word)).allMatches(words.join('')).length;
            obj[word] = count;    
        }
        int windowSize = words.fold(0, (sum, str) => sum + str.length);
        int size = s.length - windowSize;
        for (int i = 0; i <= size; ++i)
        {
            String str = s.substring(i, i + windowSize);
            Map<String, int> map = {};
            bool flag = true;
            for (var word in words)
            {
                int count = RegExp(RegExp.escape(word)).allMatches(str).length;
                map[word] = count;
            }

            for (var word in words)
            {
                if (obj[word] != map[word])
                {
                    flag = false;
                    break;
                }
            }
            
            if (flag)   startOfSubString.add(i);
        }
        return startOfSubString;
  }
}

main()
{
    var sol = Solution();
    sol.findSubstring("ababaab", ["ab","ba","ba"]).forEach((n) => print(n));   
}