"""
    Given a sorted array nums,
    remove the duplicates in-place such that each element appear only once and return the new length.

    Do not allocate extra space for another array,
    you must do this by modifying the input array in-place with O(1) extra memory.
"""

class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        nums[:] = list(set(nums))
        nums.sort()
        return len(nums)
        

    def test(self):
        tests = [
            [1,1,1,2,3,4,5],
            [1,1,1,1,2,3,4,5,1],
            [1,1,1,2,2,2,3,3,4,1],
            [1,1,1]
        ]
        for x in tests:
            print(u"sample list: {0}, len: {1}".format(x, len(x)))
            result = self.removeDuplicates(x)
            print(u'after remove duplucates: {0}, len: {1}'.format(x, len(x)))
            print(u'=======================================================')

solution = Solution()
solution.test()
