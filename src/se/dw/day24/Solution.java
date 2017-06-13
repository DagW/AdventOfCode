package se.dw.day24;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

/**
 * My solution to the day24 problem
 * http://adventofcode.com/2016/day/24
 * <p>
 * Visit all numbered points on the map, using the lowest possible amount of moves
 * <p>
 * it seems to be a good use case for one of these?
 * https://en.wikipedia.org/wiki/Travelling_salesman_problem
 * https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
 * https://en.wikipedia.org/wiki/Breadth-first_search
 * <p>
 * Im guessing:
 * 1st: calculate all distances from points of interest
 * 2nd: Generate all paths possible (Kind of brute force..)
 * 3rd: Calc distances for all paths
 * <p>
 * Permutation from here:
 * https://stackoverflow.com/questions/4240080/generating-all-permutations-of-a-given-string
 */
public class Solution {

    public Solution(String input) {

        //Create the map
        RobotMap map = new RobotMap(input);

        //Show the map
        //map.printMap(null);

        char[] pointsOfInterest = map.getPointsOfInterest().toCharArray();

        Map<Character, Map<Character, Integer>> m = new HashMap<>();
        for (int i = 0; i < pointsOfInterest.length; i++) {

            //To try the different algorithms
            //m.put(pointsOfInterest[i], map.getDistancesBFS(pointsOfInterest[i]));
            m.put(pointsOfInterest[i], map.getDistancesDijkstra(pointsOfInterest[i]));

        }

        //Get all permutations starting with 0
        permutation("0", map.getPointsOfInterest().replace("0", ""));

        int bestDistance = Integer.MAX_VALUE;
        String bestPermutation = "";

        for (String permutation : permutations) {

            int distance = 0;
            char[] chars = permutation.toCharArray();

            for (int i = 0; i < chars.length - 1; i++) {
                distance += m.get(chars[i]).get(chars[i + 1]);
            }

            System.out.println("" + permutation + " => " + distance);

            if (distance < bestDistance) {
                bestDistance = distance;
                bestPermutation = permutation;
            }

        }

        System.out.println("Solution: " + bestPermutation + ", distance=" + bestDistance);

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
                permutation(prefix + str.charAt(i), str.substring(0, i) + str.substring(i + 1, n));
        }
    }

    public static void main(String[] args) throws IOException {

        Path path = Paths.get("res/day24.txt");
        String input = new String(Files.readAllBytes(path));
        new Solution(input);

    }

}