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
            
          

""" TEST """
list1 = [1,1,1,2,3,4,5]
list2 = [1,1,1,1,2,3,4,5,1]

sol = Solution()
print('test case => ', list1)
result = sol.removeDuplicates(list1)
print('result => ', list1, ' len ', result)

print('test case => ', list2)
result = sol.removeDuplicates(list2)
print('result => ', list2, ' len ', result)
