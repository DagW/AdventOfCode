package se.dw.day24;

import java.io.IOException;
import java.util.ArrayList;

/**
 * My solution to the day24 problem - Not complete!
 * http://adventofcode.com/2016/day/24
 * <p>
 * Visit all numbered points on the map, using the lowest possible amount of moves
 * <p>
 * it seems to be a good use case for one of these?
 * https://en.wikipedia.org/wiki/Travelling_salesman_problem
 * https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
 * https://en.wikipedia.org/wiki/Breadth-first_search
 *
 * Im guessing:
 * 1st: calculate all distances from points of interest
 * 2nd: Generate in-code graph
 * 3rd: Navigate
 *
 * Permutation from here:
 * https://stackoverflow.com/questions/4240080/generating-all-permutations-of-a-given-string
 *
 */
public class Solution {

    static String exampleMap = "###########\n" +
            "#0.1.....2#\n" +
            "#.#######.#\n" +
            "#4.......3#\n" +
            "###########";

    public class RobotMap {

        private int[] startingPosition;
        private int[] currentPosition;
        private char[][] map2d;
        private ArrayList<int[]> pointsOfInterest = new ArrayList<>();

        public RobotMap(String input) {

            //Create the map and find all numbers
            map2d = new char[input.split("\n").length][input.split("\n")[0].length()];

            String perm = "";
            for (int y = 0; y < input.split("\n").length; y++) {

                char[] chars = input.split("\n")[y].toCharArray();
                map2d[y] = chars;

                for (int x = 0; x < chars.length; x++) {
                    if (Character.isDigit(chars[x])) {
                        if (chars[x] == '0') {
                            startingPosition = new int[]{y, x};
                        } else {
                            pointsOfInterest.add(new int[]{y, x});
                        }

                        perm += ""+chars[x];
                    }
                }
            }

            //We start at the 0
            currentPosition = startingPosition;

            //Generate all distances
            permutation("",perm);

            for (String p : permutations)
                System.out.println(p);

            //TODO Get cost for each permutation

        }

    }


    /*
    https://stackoverflow.com/questions/4240080/generating-all-permutations-of-a-given-string
     */
    ArrayList<String> permutations = new ArrayList<>();
    private void permutation(String prefix, String str) {
        int n = str.length();
        if (n == 0)
            permutations.add(prefix);
        else {
            for (int i = 0; i < n; i++)
                permutation(prefix + str.charAt(i), str.substring(0, i) + str.substring(i+1, n));
        }
    }

    public Solution(String input) {

        //Create the map
        RobotMap map = new RobotMap(input);
        //Get weight from 0 to 1


    }

    public static void main(String[] args) throws IOException {

        //Path path = Paths.get("res/day24.txt");
        //String input = new String(Files.readAllBytes(path));
        new Solution(exampleMap);

    }

}