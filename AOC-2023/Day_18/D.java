import java.io.File;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.Scanner;

class Pair {
    Long x;
    Long y;

    public Pair(Long x, Long y) {
        this.x = x;
        this.y = y;
    }
}

class D {
    public static void main(String[] args) {
        try {
            Scanner sc = new Scanner(new File("I.txt"));
            HashMap<String, Pair> map = new HashMap<>();
            map.put("U", new Pair(-1L, 0L));
            map.put("R", new Pair(0L, 1L));
            map.put("D", new Pair(1L, 0L));
            map.put("L", new Pair(0L, -1L));
            map.put("3", new Pair(-1L, 0L));
            map.put("0", new Pair(0L, 1L));
            map.put("1", new Pair(1L, 0L));
            map.put("2", new Pair(0L, -1L));

            LinkedList<Pair> points = new LinkedList<>();
            Long B = 0L;

            Long currX = 0L;
            Long currY = 0L;

            while (sc.hasNextLine()) {
                String data = sc.nextLine().split(" ")[2].substring(2);
                data = data.substring(0, data.length() - 1);

                Pair p = map.get(data.substring(data.length() - 1, data.length()));
                // System.out.println(data.substring(data.length() - 1, data.length()));
                data = data.substring(0, data.length() - 1);
                int num = Integer.parseInt(data, 16);
                B += num;
                points.add(new Pair(currX + num * p.x, currY + num * p.y));
                currX = currX + num * p.x;
                currY = currY + num * p.y;
                // System.out.println(data + " val: " + num + " dir: " + p.x + " , " + p.y);

                // B += num;
                /*
                 * PART ONE
                 * 
                 * int num = Integer.parseInt(data[1]);
                 * B += num;
                 * Pair p = map.get(data[0]);
                 * points.add(new Pair(currX + num * p.x, currY + num * p.y));
                 * currX = currX + num * p.x;
                 * currY = currY + num * p.y;
                 */

            }

            points.add(0, points.getLast());
            points.addLast(points.get(1));
            Long A = 0L;
            for (int i = 1; i < points.size() - 1; i++) {
                A += points.get(i).x * (points.get(i - 1).y - points.get(i + 1).y);
            }

            A = Math.abs(A) / 2;
            Long smallI = A - B / 2 + 1;
            smallI += B;
            System.out.println(smallI);

        } catch (Exception e) {
            System.err.println(e);
        }
    }
}