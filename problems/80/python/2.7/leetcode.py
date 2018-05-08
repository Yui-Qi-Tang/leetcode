"""
problem 80:
        Given a sorted array nums, 
        remove the duplicates in-place such that duplicates appeared at most twice and return the new length.

        Do not allocate extra space for another array, 
        you must do this by modifying the input array in-place with O(1) extra memory.
"""

class Solution:
    """
    def removeDuplicates(self, nums):
    
        #:type nums: List[int]
        #:rtype: int
        
        find = lambda searchList, elem: [[i for i, x in enumerate(searchList) if x == e] for e in elem]

        for x in find(nums, set(nums)):
            if len(x) > 2:
                del nums[x[0]]
        return len(nums)
    """

    def removeDuplicates(self, nums):
        for index, v in enumerate(nums):
            i = 0
            for y in nums[index+1:]:
                if v !=y:
                    break
                if v == y and i < 2:
                    i = i + 1
                if i >= 2:
                    nums.remove(y)
        nums[:] = nums
        return len(nums)
            
    def test(self):
        tests = [
            [1,1,1,2,3,4,5],
            [1,1,1,1,2,3,4,5,1],
        ]
        for x in tests:
            print(u"sample list: {0}, len: {1}".format(x, len(x)))
            result = self.removeDuplicates(x)
            print(u'after remove duplucates: {0}, len: {1}'.format(x, len(x)))
            print(u'=======================================================')

solution = Solution()
solution.test()