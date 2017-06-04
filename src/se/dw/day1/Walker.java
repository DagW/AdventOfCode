package se.dw.day1;

public class Walker {

    final int NORTH = 0, SOUTH = 2, EAST = 1, WEST = 3;
    int orientation = NORTH;

    //Coordinates on the city grid
    private int[] coordinates = new int[]{0, 0};

    public void moveLeft(int distance) {
        //First we turn left
        orientation -= 1;
        orientation = orientation % 4;
        if (orientation < NORTH)
            orientation = WEST;
        //Then we move
        move(distance);
    }

    public void moveRight(int distance) {
        //First we turn right
        orientation += 1;
        orientation = orientation % 4;
        //Then we move
        move(distance);
    }

    private void move(int distance) {

        //Vertical movement
        if (orientation == NORTH)
            coordinates[0] += distance;
        else if (orientation == SOUTH)
            coordinates[0] -= distance;
            //Horizontal movement
        else if (orientation == EAST)
            coordinates[1] += distance;
        else if (orientation == WEST)
            coordinates[1] -= distance;

    }

    public int getCityBlockDistance() {
        // https://en.wikipedia.org/wiki/Taxicab_geometry
        return Math.abs(coordinates[0]) + Math.abs(coordinates[1]);
    }

}