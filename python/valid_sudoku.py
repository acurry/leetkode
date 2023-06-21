def isValidSudoku(board):
	rows = []
	cols = []
	transposed_board = list(map(list, zip(*board)))
	for i in range(9):
		rows.append([x for x in board[i] if x is not '.'])
	for i in range(9):
		rows.append([x for x in transposed_board[i] if x is not '.')
	
	

# [["5","3",".",".","7",".",".",".","."]
#, ["6",".",".","1","9","5",".",".","."]
#, [".","9","8",".",".",".",".","6","."]
#, ["8",".",".",".","6",".",".",".","3"]
#, ["4",".",".","8",".","3",".",".","1"]
#, ["7",".",".",".","2",".",".",".","6"]
#, [".","6",".",".",".",".","2","8","."]
#, [".",".",".","4","1","9",".",".","5"]
#, [".",".",".",".","8",".",".","7","9"]]