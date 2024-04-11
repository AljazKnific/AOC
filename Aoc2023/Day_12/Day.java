import java.io.File;
import java.util.LinkedList;
import java.util.Scanner;

public class Day {

    public static boolean onlyDots(String data, int x, int end) {
        System.out.println("Data: " + data + " x: " + x + " end: " + end + " len: " +
                data.length());

        if (end > data.length() || x >= data.length()) {

            return true;
        }

        for (int i = x; i < end; i++) {

            if (data.charAt(i) != '.' && data.charAt(i) != '?') {

                return false;

            }
        }

        return true;
    }

    public static int rec(LinkedList<Integer> list, int id, String data) {
        int sum = 0;
        System.out.println(data + " ID: " + list.get(id) + " i: " + id + " data.len: " + data.length());

        if (list.size() == id && onlyDots(data, 0, data.length())) {
            return 1;
        }

        // we are at the end of the string
        /*
         * 
         * if (id == list.size() || data.length() < list.get(id)) {
         * return 0;
         * }
         */

        // searching for substrings inside the data
        for (int i = 0; i < data.length(); i++) {
            boolean bo = true;
            if (i + list.get(id) > data.length()) {
                break;
            }

            for (int j = 0; j < list.get(id); j++) {
                if (data.charAt(i + j) == '.') {
                    // preskocimo piko in nadaljujemo od tam naprej
                    bo = false;
                    i += j;
                    break;
                }
            }

            if (bo && onlyDots(data, i + list.get(id), i + list.get(id) + 1)) {
                System.out.println("BOMM : i + 2 -> " + (i + 2) + " d.l -> " + data.length());
                sum += rec(list, id + 1, data.substring(i + 2, data.length() - 1));
            }
        }

        return sum;
    }

    public static void main(String[] args) {
        try {
            Scanner sc = new Scanner(new File("Input.txt"));
            int sum = 0;

            while (sc.hasNextLine()) {
                String[] array = sc.nextLine().split(" ");
                LinkedList<Integer> list = new LinkedList<>();
                String[] nums = array[1].split(",");
                for (int i = 0; i < nums.length; i++) {
                    list.add(Integer.parseInt(nums[i]));
                }

                sum += rec(list, 0, array[0]);

            }

            System.out.println("SUM: " + sum);

            sc.close();

        } catch (Exception e) {
            System.err.println(e);
        }
    }
}
