package se.dw.day1;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class SolutionTest {

    private static final String input = "R3, L5, R2, L2, R1, L3, R1, R3, L4, R3, L1, L1, R1, L3, R2, L3, L2, R1, " +
            "R1, L1, R4, L1, L4, R3, L2, L2, R1, L1, R5, R4, R2, L5, L2, R5, R5, L2, R3, R1, R1, L3, R1, L4, L4, " +
            "L190, L5, L2, R4, L5, R4, R5, L4, R1, R2, L5, R50, L2, R1, R73, R1, L2, R191, R2, L4, R1, L5, L5, " +
            "R5, L3, L5, L4, R4, R5, L4, R4, R4, R5, L2, L5, R3, L4, L4, L5, R2, R2, R2, R4, L3, R4, R5, L3, R5, " +
            "L2, R3, L1, R2, R2, L3, L1, R5, L3, L5, R2, R4, R1, L1, L5, R3, R2, L3, L4, L5, L1, R3, L5, L2, R2, " +
            "L3, L4, L1, R1, R4, R2, R2, R4, R2, R2, L3, L3, L4, R4, L4, L4, R1, L4, L4, R1, L2, R5, R2, R3, R3, " +
            "L2, L5, R3, L3, R5, L2, R3, R2, L4, L3, L1, R2, L2, L3, L5, R3, L1, L3, L4, L3";

    Solution solution;

    @org.junit.Before
    public void setUp() throws Exception {
        solution = new Solution(input);
    }

    @Test
    public void testSolution() {
        assertEquals(solution.getWalker().getCityBlockDistance(), 291);
    }

    @Test
    public void testWalker() {
        Walker walker = new Walker();
        walker.moveLeft(5);
        assertEquals(walker.getCityBlockDistance(), 5);
    }

    @org.junit.After
    public void tearDown() throws Exception {

    }

}