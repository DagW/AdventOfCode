package se.dw.aoc2016.day24;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class SolutionTest {

    static String exampleMap = "###########\n" +
            "#0.1.....2#\n" +
            "#.#######.#\n" +
            "#4.......3#\n" +
            "###########";

    RobotMap map;

    @org.junit.Before
    public void setUp() throws Exception {

        map = new RobotMap(exampleMap);

    }

    @Test
    public void testRobotMap() {

        //It should contain the numbers but no walls or dots
        assertEquals(map.getPointsOfInterest().contains("0"), true);
        assertEquals(map.getPointsOfInterest().contains("1"), true);
        assertEquals(map.getPointsOfInterest().contains("2"), true);
        assertEquals(map.getPointsOfInterest().contains("3"), true);
        assertEquals(map.getPointsOfInterest().contains("4"), true);
        assertEquals(map.getPointsOfInterest().contains("."), false);
        assertEquals(map.getPointsOfInterest().contains("#"), false);

        assertEquals(map.getDistancesBFS('0').size(), map.getPointsOfInterest().length());

        /*
        From the example
        0 to 4 (2 steps)
        4 to 1 (4 steps; it can't move diagonally)
        1 to 2 (6 steps)
        2 to 3 (2 steps)
         */
        assertEquals((int) map.getDistancesBFS('0').get('0'), 0);
        assertEquals((int) map.getDistancesBFS('0').get('4'), 2);
        assertEquals((int) map.getDistancesBFS('4').get('1'), 4);
        assertEquals((int) map.getDistancesBFS('1').get('2'), 6);
        assertEquals((int) map.getDistancesBFS('2').get('3'), 2);
        assertEquals((int) map.getDistancesDijkstra('0').get('0'), 0);
        assertEquals((int) map.getDistancesDijkstra('0').get('4'), 2);
        assertEquals((int) map.getDistancesDijkstra('4').get('1'), 4);
        assertEquals((int) map.getDistancesDijkstra('1').get('2'), 6);
        assertEquals((int) map.getDistancesDijkstra('2').get('3'), 2);

    }

    @org.junit.After
    public void tearDown() throws Exception {

    }

}