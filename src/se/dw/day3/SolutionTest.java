package se.dw.day3;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class SolutionTest {

    static String validTriangle = "100 100 100";
    static String invalidTriange = "1 90 100";

    @org.junit.Before
    public void setUp() throws Exception {
    }

    @Test
    public void testSolution() {
        assertEquals(new Solution(validTriangle).validTriangleCount, 1);
        assertEquals(new Solution(invalidTriange).validTriangleCount, 0);
    }

    @org.junit.After
    public void tearDown() throws Exception {

    }

}