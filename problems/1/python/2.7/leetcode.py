class Solution(object):
    """
        Hint: 1.there are two elements as same as target extractly!
              2.do not use the same elements twice
    """
    def twoSum(self, nums, target):
        """
        Just use list computing(slowest)
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        # if not nums: return []


        for i in range(len(nums)):
            d = target - nums[i] # use Hint 1
            if nums.count(d) > 1: # gerantee the d is in list!!
                j = nums.index(d)
                if i!=j: # check Hint 2
                    return [i, j]
            
    def twoSum1(self, nums, target):
        """
        Use the 'try except' (middle)
        I think the line 14 that used count function of list to check element numbers 
        This version use the 'try except' to check element numbers of the list
        (Is that correct?) 
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        # if not nums: return []


        for i in range(len(nums)):
            d = target - nums[i]
            try:
                j = nums.index(d) # do first without check the the d is whether in list or not
                if i!=j:
                    return [i, j]
            except ValueError:
                pass

    def twoSum2(self, nums, target):
        """
        Use dict to store data(fastest)
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        # if not nums: return []

        diff_records = {}
        for k, v in enumerate(nums):
            d = target - v # Hint 1
            if d in diff_records:
                return [diff_records[d], k]
            diff_records[v] = k # save v and k for next round
               

    def test(self):
        """
            Given nums = [2, 7, 11, 15], target = 9,

            Because nums[0] + nums[1] = 2 + 7 = 9,
            return [0, 1].
        """
        target = 6
        test = [3,2,4]
        result = self.twoSum2(test, target)
        print(result)
        
sol = Solution()
sol.test()