package day08

import (
	"testing"
)

func Test_Layers(t *testing.T) {
	if 1224 != solvePart1() {
		t.Errorf("wrong answer ")
	}

	render()
	/*
		#### ###  #### #  # ###
		#    #  #    # #  # #  #
		###  ###    #  #  # #  #
		#    #  #  #   #  # ###
		#    #  # #    #  # # #
		#### ###  ####  ##  #  #
	*/
}
