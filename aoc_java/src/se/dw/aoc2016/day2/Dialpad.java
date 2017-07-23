package se.dw.aoc2016.day2;

public class Dialpad {

    private static final int X = 1, Y = 0;

    // representing a dialpad
    int[][] dials = {
            {1, 2, 3},
            {4, 5, 6},
            {7, 8, 9}
    };
    // current coordinate on the dialpad, starting at 5 (x,y) -> (1,1)
    int[] dialcoordinate = {1, 1};

    public int getCurrentDial() {
        return dials[dialcoordinate[Y]][dialcoordinate[X]];
    }

    public Dialpad parseInstruction(String instruction) {

        for (char ch : instruction.toCharArray()) {
            switch (ch) {
                case 'U':
                    moveUp();
                    break;
                case 'D':
                    moveDown();
                    break;
                case 'L':
                    moveLeft();
                    break;
                case 'R':
                    moveRight();
                    break;
            }
        }

        return this;
    }

    public void moveUp() {
        //If the move is within the dialpad
        if (dialcoordinate[Y] > 0)
            dialcoordinate[Y]--;
    }

    public void moveDown() {
        if (dialcoordinate[Y] < 2)
            dialcoordinate[Y]++;
    }

    public void moveLeft() {
        if (dialcoordinate[X] > 0)
            dialcoordinate[X]--;
    }

    public void moveRight() {
        if (dialcoordinate[X] < 2)
            dialcoordinate[X]++;
    }

}