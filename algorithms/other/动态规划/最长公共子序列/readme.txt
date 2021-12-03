ints1 := []int{1, 3, 5, 9, 1}
ints2 := []int{1, 4, 9, 10}

dp(i,j): 是nums1的前1个元素,nums2前j个元素

dp(2,3)=1
    1,3
    1,4,9
dp(4,4)=2
    1, 3, 5, 9
    1, 4, 9, 10


最终: dp(len(nums1),len(num2))


如果: nums1[i-1]=nums2[j-1]
dp(i,j)=dp(i-1,j-1)+1

nums1[i-1]!=nums2[j-1]
dp(i,j)=max(dp(i-1,j),dp(i,j-1))










