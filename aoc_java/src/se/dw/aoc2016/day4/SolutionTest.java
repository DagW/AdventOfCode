package se.dw.aoc2016.day4;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class SolutionTest {

    public final static String validInput1 = "aaaaa-bbb-z-y-x-123[abxyz]";
    public final static String validInput2 = "a-b-c-d-e-f-g-h-987[abcde]";
    public final static String validInput3 = "not-a-real-room-404[oarel]";
    public final static String invalidInput1 = "totally-real-room-200[decoy]";

    @org.junit.Before
    public void setUp() throws Exception {
    }

    @Test
    public void testCalculateChecksum() {
        assertEquals(new Room(validInput1).getChecksum(), "abxyz");
        assertEquals(new Room(validInput2).getChecksum(), "abcde");
        assertEquals(new Room(validInput3).getChecksum(), "oarel");
        assertEquals(new Room(invalidInput1).getChecksum(), "decoy");
    }

    @Test
    public void testGetRoomId() {
        assertEquals(new Room(validInput1).getRoomId(), 123);
        assertEquals(new Room(validInput2).getRoomId(), 987);
        assertEquals(new Room(validInput3).getRoomId(), 404);
        assertEquals(new Room(invalidInput1).getRoomId(), 200);
    }

    @org.junit.After
    public void tearDown() throws Exception {

    }

}