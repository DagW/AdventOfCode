package se.dw.day1;

/**
 * My solution to the day1 problem
 * http://adventofcode.com/2016/day/1
 * https://en.wikipedia.org/wiki/Taxicab_geometry
 */
public class Solution {

    private static final String input = "R3, L5, R2, L2, R1, L3, R1, R3, L4, R3, L1, L1, R1, L3, R2, L3, L2, R1, " +
            "R1, L1, R4, L1, L4, R3, L2, L2, R1, L1, R5, R4, R2, L5, L2, R5, R5, L2, R3, R1, R1, L3, R1, L4, L4, " +
            "L190, L5, L2, R4, L5, R4, R5, L4, R1, R2, L5, R50, L2, R1, R73, R1, L2, R191, R2, L4, R1, L5, L5, " +
            "R5, L3, L5, L4, R4, R5, L4, R4, R4, R5, L2, L5, R3, L4, L4, L5, R2, R2, R2, R4, L3, R4, R5, L3, R5, " +
            "L2, R3, L1, R2, R2, L3, L1, R5, L3, L5, R2, R4, R1, L1, L5, R3, R2, L3, L4, L5, L1, R3, L5, L2, R2, " +
            "L3, L4, L1, R1, R4, R2, R2, R4, R2, R2, L3, L3, L4, R4, L4, L4, R1, L4, L4, R1, L2, R5, R2, R3, R3, " +
            "L2, L5, R3, L3, R5, L2, R3, R2, L4, L3, L1, R2, L2, L3, L5, R3, L1, L3, L4, L3";

    WalkerState walkerState = new WalkerState();

    public Solution(String input) {

        /*
        Example input - R3, L5
        We start at location[0,0], facing north
        walking Right 3 blocks to location[0,3], facing east
        we then walk Left 5 blocks to location[5,3], facing north
        The "City block distance" back would then be 8.
         */

        String[] instructions = input.split(", ");
        for (String instruction : instructions) {
            walk(instruction);
        }

        System.out.println("Solution: " + walkerState.getCityBlockDistance());
    }

    private void walk(String instruction) {

        char turnDirection = instruction.charAt(0);
        int distance = Integer.parseInt(instruction.substring(1));

        switch (turnDirection) {
            case 'L':
                walkerState.moveLeft(distance);
                break;
            case 'R':
                walkerState.moveRight(distance);
                break;
        }

    }

    public class WalkerState {

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

            if (orientation == NORTH)
                coordinates[0] += distance;
            else if (orientation == SOUTH)
                coordinates[0] -= distance;
            else if (orientation == EAST)
                coordinates[1] += distance;
            else if (orientation == WEST)
                coordinates[1] -= distance;

        }

        public int getCityBlockDistance() {
            return Math.abs(coordinates[0]) + Math.abs(coordinates[1]);
        }

    }

    public static void main(String[] args) {
        new Solution(input);
    }
}