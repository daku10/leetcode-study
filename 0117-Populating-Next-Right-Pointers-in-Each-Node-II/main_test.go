package main

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	testcases := []struct {
		arg  *Node
		want *Node
	}{
		// {
		// 	&Node{
		// 		Val: 1,
		// 		Left: &Node{
		// 			Val: 2,
		// 			Left: &Node{
		// 				Val: 4,
		// 			},
		// 			Right: &Node{
		// 				Val: 5,
		// 			},
		// 		},
		// 		Right: &Node{
		// 			Val: 3,
		// 			Left: &Node{
		// 				Val: 6,
		// 			},
		// 			Right: &Node{
		// 				Val: 7,
		// 			},
		// 		},
		// 	},
		// 	func() *Node {
		// 		one := &Node{
		// 			Val: 1,
		// 		}
		// 		two := &Node{
		// 			Val: 2,
		// 		}
		// 		three := &Node{
		// 			Val: 3,
		// 		}
		// 		four := &Node{
		// 			Val: 4,
		// 		}
		// 		five := &Node{
		// 			Val: 5,
		// 		}
		// 		six := &Node{
		// 			Val: 6,
		// 		}
		// 		seven := &Node{
		// 			Val: 7,
		// 		}
		// 		one.Left = two
		// 		one.Right = three
		// 		two.Left = four
		// 		two.Right = five
		// 		three.Left = six
		// 		three.Right = seven
		// 		two.Next = three
		// 		four.Next = five
		// 		five.Next = six
		// 		six.Next = seven
		// 		return one
		// 	}(),
		// },
		// {
		// 	nil,
		// 	nil,
		// },
		// {
		// 	&Node{
		// 		Val: 1,
		// 		Left: &Node{
		// 			Val: 2,
		// 			Left: &Node{
		// 				Val: 4,
		// 			},
		// 			Right: &Node{
		// 				Val: 5,
		// 			},
		// 		},
		// 		Right: &Node{
		// 			Val: 3,
		// 			Right: &Node{
		// 				Val: 7,
		// 			},
		// 		},
		// 	},
		// 	func() *Node {
		// 		one := &Node{
		// 			Val: 1,
		// 		}
		// 		two := &Node{
		// 			Val: 2,
		// 		}
		// 		three := &Node{
		// 			Val: 3,
		// 		}
		// 		four := &Node{
		// 			Val: 4,
		// 		}
		// 		five := &Node{
		// 			Val: 5,
		// 		}
		// 		seven := &Node{
		// 			Val: 7,
		// 		}
		// 		one.Left = two
		// 		one.Right = three
		// 		two.Left = four
		// 		two.Right = five
		// 		three.Right = seven
		// 		two.Next = three
		// 		four.Next = five
		// 		five.Next = seven
		// 		return one
		// 	}(),
		// },
		// [2,1,3,0,7,9,1,2,null,1,0,null,null,8,8,null,null,null,null,7]
		{
			&Node{
				Val: 2,
				Left: &Node{
					Val: 1,
					Left: &Node{
						Val: 0,
						Left: &Node{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &Node{
						Val: 7,
						Left: &Node{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: &Node{
							Val: 0,
							Left: &Node{
								Val: 7,
							},
						},
					},
				},
				Right: &Node{
					Val: 3,
					Left: &Node{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
					Right: &Node{
						Val: 1,
						Left: &Node{
							Val: 8,
						},
						Right: &Node{
							Val: 8,
						},
					},
				},
			},
			func() *Node {
				n1 := &Node{
					Val: 2,
				}
				n2 := &Node{
					Val: 1,
				}
				n3 := &Node{
					Val: 3,
				}
				n4 := &Node{
					Val: 0,
				}
				n5 := &Node{
					Val: 7,
				}
				n6 := &Node{
					Val: 9,
				}
				n7 := &Node{
					Val: 1,
				}
				n8 := &Node{
					Val: 2,
				}
				n9 := &Node{
					Val: 1,
				}
				n10 := &Node{
					Val: 0,
				}
				n11 := &Node{
					Val: 8,
				}
				n12 := &Node{
					Val: 8,
				}
				n13 := &Node{
					Val: 7,
				}
				n1.Left = n2
				n1.Right = n3
				n2.Left = n4
				n2.Right = n5
				n3.Left = n6
				n3.Right = n7
				n4.Left = n8
				n4.Right = nil
				n5.Left = n9
				n5.Right = n10
				n6.Left = nil
				n6.Right = nil
				n7.Left = n11
				n7.Right = n12
				n10.Left = n13
				n2.Next = n3
				n4.Next = n5
				n5.Next = n6
				n6.Next = n7
				n8.Next = n9
				n9.Next = n10
				n10.Next = n11
				n11.Next = n12
				return n1
			}(),
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := connect(tc.arg)
			if !cmpTree(got, tc.want) {
				t.Errorf("got: %v want: %v", got, tc.want)
				printTree(got)
				fmt.Println()
				printTree(tc.want)
				fmt.Println()
			}
		})
	}
}

func printTree(node *Node) {
	if node == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(node.Val)
	printTree(node.Left)
	printTree(node.Right)
}

func cmpTree(x, y *Node) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		fmt.Println("x || y nil")
		return false
	}
	if x.Next != nil && y.Next == nil {
		fmt.Println("x.Next is not nil || y.Next is nil")
		return false
	}
	if x.Next == nil && y.Next != nil {
		fmt.Println("x.Next is nil || y.Next is not nil")
		return false
	}
	if x.Val != y.Val {
		fmt.Println("x.Val is not y.Val")
		return false
	}
	if x.Next != nil && x.Next.Val != y.Next.Val {
		fmt.Println("x.Next.Val is not y.Next.Val")
		fmt.Println(x.Val, y.Val)
		fmt.Println(x.Next.Val, y.Next.Val)
		return false
	}
	return cmpTree(x.Left, y.Left) && cmpTree(x.Right, y.Right)
}
