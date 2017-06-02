package se.dw.day2;

import org.junit.Test;
import static org.junit.Assert.assertEquals;

public class SolutionTest {

    public final static String input1 = "ULDRULRDULDRD";
    public final static String input2 = "UUUUUUU";
    public final static String input3 = "RURURU";
    public final static String input4 = "DLDLDLDL";

    @org.junit.Before
    public void setUp() throws Exception {
    }

    @Test
    public void testSolution() {
        assertEquals(new Solution(input1).getDialpad().getCurrentDial(), 8);
        assertEquals(new Solution(input2).getDialpad().getCurrentDial(), 2);
        assertEquals(new Solution(input3).getDialpad().getCurrentDial(), 3);
        assertEquals(new Solution(input4).getDialpad().getCurrentDial(), 7);
    }

    @org.junit.After
    public void tearDown() throws Exception {

    }

}