package se.dw.day24;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class RobotMap {

    private char[][] map2d;
    private String pointsOfInterest = "";

    public RobotMap(String input) {

        //Create the map and find all numbers
        map2d = new char[input.split("\n").length][input.split("\n")[0].length()];

        for (int y = 0; y < input.split("\n").length; y++) {

            char[] chars = input.split("\n")[y].toCharArray();
            map2d[y] = chars;

            for (int x = 0; x < chars.length; x++) {

                if (Character.isDigit(chars[x])) {
                    pointsOfInterest += "" + chars[x];
                }

            }

        }

    }

    public String getPointsOfInterest() {
        return pointsOfInterest;
    }


    private Position getPositionOf(char character) {

        for (int y = 0; y < map2d.length; y++) {
            for (int x = 0; x < map2d[y].length; x++) {

                if (map2d[y][x] == character) {
                    return new Position(x, y);
                }

            }
        }

        return null;
    }

    void printMap(Position position) {

        for (int i = 0; i < map2d.length; i++) {
            for (int j = 0; j < map2d[i].length; j++) {

                if (position != null && i == position.getY() && j == position.getX()) {
                    System.out.print("X");
                } else {
                    System.out.print(map2d[i][j]);
                }

            }

            System.out.println();
        }

    }


    /*
    BFS search for all other chars
    https://en.wikipedia.org/wiki/Breadth-first_search#Pseudocode
     */
    public Map<Character, Integer> getDistancesFrom(char startchar) {

        Map<Character, Integer> distanceMap = new HashMap<>();

        Position currentPosition = getPositionOf(startchar);

        ArrayList<Position> visitedPositions = new ArrayList<>();
        ArrayList<Position> positionsToVisit = new ArrayList<>();
        positionsToVisit.add(currentPosition);

        while (!positionsToVisit.isEmpty()) {

            //Move to the next node
            currentPosition = positionsToVisit.remove(0);

            if (!visitedPositions.contains(currentPosition)) {

                //System.out.println("Testing "+currentPosition);

                if (pointsOfInterest.contains("" + currentPosition.getChar())) {
                    //System.out.println(start.getChar() + " -> " + currentPosition.getChar() + " is " + currentPosition.getDistance());

                    distanceMap.put(currentPosition.getChar(), currentPosition.getDistance());

                    if (distanceMap.size() == pointsOfInterest.length()) {
                        //We found all distances
                        return distanceMap;
                    }

                }

                Position[] adjacentPositions = new Position[]{
                        new Position(currentPosition.getX(), currentPosition.getY() + 1, currentPosition.getDistance() + 1), //Above
                        new Position(currentPosition.getX(), currentPosition.getY() - 1, currentPosition.getDistance() + 1), //Below
                        new Position(currentPosition.getX() + 1, currentPosition.getY(), currentPosition.getDistance() + 1), //Left
                        new Position(currentPosition.getX() - 1, currentPosition.getY(), currentPosition.getDistance() + 1) //Right
                };

                for (Position position : adjacentPositions) {
                    if (!visitedPositions.contains(position) && position.isWalkable()) {

                        positionsToVisit.add(position);

                    }
                }

                visitedPositions.add(currentPosition);

            }
        }

        return distanceMap;
    }


    private class Position {
        int x, y, distance = 0;

        public Position(int x, int y) {
            this.x = x;
            this.y = y;
        }

        @Override
        public boolean equals(Object obj) {
            Position p2 = (Position) obj;
            if (this.x == p2.getX() && this.y == p2.getY()) {
                return true;
            }
            return false;
        }

        public Position(int x, int y, int distance) {
            this.x = x;
            this.y = y;
            this.distance = distance;
        }

        public int getDistance() {
            return distance;
        }

        @Override
        public String toString() {
            return "Position(" + x + "," + y + ") d=" + distance + " [" + getChar() + "]";
        }

        public boolean isWalkable() {

            //Check that we are inside map bounds
            if (this.x < 0 || this.x >= map2d[0].length - 1
                    || this.y < 0 || this.y >= map2d.length - 1) {
                return false;
            }
            //Check that the position is not a wall
            if (map2d[this.y][this.x] == '#') {
                return false;
            }

            return true;
        }

        public char getChar() {
            return map2d[this.y][this.x];
        }

        public int getY() {
            return y;
        }

        public int getX() {
            return x;
        }
    }

}