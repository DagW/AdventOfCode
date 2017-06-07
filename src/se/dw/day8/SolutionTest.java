package se.dw.day8;

import org.junit.Test;
import static org.junit.Assert.assertEquals;

public class SolutionTest {


    @org.junit.Before
    public void setUp() throws Exception {
    }

    @Test
    public void testScreen(){

        Screen screen = new Screen(5,10);
        assertEquals(screen.getLitPixels(), 0);

        screen.parseCommand("rect 5x1");
        screen.parseCommand("rotate column y=0 by 2");
        screen.parseCommand("rotate row y=0 by 2");
        screen.parseCommand("rect 5x5");

        assertEquals(screen.getLitPixels(), 25);

        screen.parseCommand("rect 5x10");
        assertEquals(screen.getLitPixels(), 50);

    }

    @org.junit.After
    public void tearDown() throws Exception {

    }

}