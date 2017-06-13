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

    private char getChar(Position position) {
        return map2d[position.getY()][position.getX()];
    }

    private boolean isWalkable(Position position) {

        //Check that we are inside map bounds
        if (position.getX() < 0 || position.getX() >= map2d[0].length - 1
                || position.getY() < 0 || position.getY() >= map2d.length - 1) {
            return false;
        }
        //Check that the position is not a wall
        if (map2d[position.getY()][position.getX()] == '#') {
            return false;
        }

        return true;
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
    public Map<Character, Integer> getDistancesBFS(char startchar) {

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

                if (pointsOfInterest.contains("" + getChar(currentPosition))) {
                    //System.out.println(start.getChar() + " -> " + currentPosition.getChar() + " is " + currentPosition.getDistance());

                    distanceMap.put(getChar(currentPosition), currentPosition.getDistance());

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
                    if (!visitedPositions.contains(position) && isWalkable(position)) {

                        positionsToVisit.add(position);

                    }
                }

                visitedPositions.add(currentPosition);

            }
        }

        return distanceMap;
    }

    /*
    Alternative way to get walk the map
    https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Pseudocode
     */
    public Map<Character, Integer> getDistancesDijkstra(char startChar) {

        Map<Character, Integer> distances = new HashMap<>();
        distances.put(startChar, 0);

        ArrayList<Position> set = new ArrayList<>();

        //Add the start position, the only one we know the distance of
        Position start = getPositionOf(startChar);
        start.setDistance(0);
        set.add(start);

        for (int y = 0; y < map2d.length; y++) {
            for (int x = 0; x < map2d[y].length; x++) {

                //Initiate all the maps positions with max distance
                Position p = new Position(x, y);
                p.setDistance(Integer.MAX_VALUE);
                //Add if its not a wall etc.
                if (isWalkable(p) && !set.contains(p)) {
                    set.add(p);
                }
            }
        }

        while (!set.isEmpty()) {

            //Get the position with the least distance
            Position currentPosition = null;
            for (Position p : set) {
                if (currentPosition == null || p.getDistance() < currentPosition.getDistance()) {
                    currentPosition = p;
                }
            }
            //And get it from the list
            currentPosition = set.remove(set.indexOf(currentPosition));

            //Get the neighbours, if they are in the list they are within bounds and not walls
            ArrayList<Position> adjacentPositions = new ArrayList<>();

            Position above = new Position(currentPosition.getX(), currentPosition.getY() + 1);
            if (set.indexOf(above) != -1) adjacentPositions.add(above);

            Position below = new Position(currentPosition.getX(), currentPosition.getY() - 1);
            if (set.indexOf(below) != -1) adjacentPositions.add(below);

            Position left = new Position(currentPosition.getX() + 1, currentPosition.getY());
            if (set.indexOf(left) != -1) adjacentPositions.add(left);

            Position right = new Position(currentPosition.getX() - 1, currentPosition.getY());
            if (set.indexOf(right) != -1) adjacentPositions.add(right);

            for (Position position : adjacentPositions) {

                position.setDistance(currentPosition.getDistance() + 1);

                if (currentPosition.getDistance() + 1 < set.get(set.indexOf(position)).getDistance()) {
                    set.set(set.indexOf(position), position);
                }

                if (Character.isDigit(getChar(position))) {
                    distances.put(getChar(position), position.distance);
                }
            }
        }

        return distances;
    }

}