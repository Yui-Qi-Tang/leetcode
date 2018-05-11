# # Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    
    def rc(self, l1, l2):
        if not l1.next and not l2.next:
            return [l1.val + l2.val]
        else:
            return self.rc(l1.next, l2.next) + [l1.val + l2.val]
        
    def get_val(self, l):
        t = []
        while(l):
            t.append(str(l.val))
            l = l.next
        return t

    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        list_val1 ="".join(self.get_val(l1))[::-1]
        list_val2 ="".join(self.get_val(l2))[::-1]
        list_sum = int(list_val1)+int(list_val2)
        result = str(list_sum)[::-1]
        return [int(x) for x in result]


