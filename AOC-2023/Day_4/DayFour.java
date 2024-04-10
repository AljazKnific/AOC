import java.io.File;
import java.util.Arrays;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.Scanner;

//IDEA:  create two sets , then do the intersections between them, result is 2^set.size() - 1

class Item {
    int val;
    int game;

    public Item(int game, int val) {
        this.val = val;
        this.game = game;
    }
}

class DayFour {

    public static void fillTheSet(HashSet<Integer> s1, HashSet<Integer> s2, String[] data) {

        boolean change = false;

        for (String string : data) {
            // System.out.println(string);

            if (string.equals(""))
                continue;

            if (string.equals("|")) {
                change = true;
                continue;
            }

            if (change) {
                s2.add(Integer.parseInt(string));
            } else {
                s1.add(Integer.parseInt(string));
            }
        }

    }

    public static void main(String[] args) {

        try {
            File inp = new File("Input.txt");

            Scanner sc = new Scanner(inp);
            LinkedList<Item> list = new LinkedList<>();

            int sum = 0;
            int maxY = 0;

            while (sc.hasNextLine()) {
                String data = sc.nextLine();
                String[] arr;
                HashSet<Integer> s1 = new HashSet<>();
                HashSet<Integer> s2 = new HashSet<>();

                // seperate game from numbers
                arr = data.split(":");

                arr = arr[1].split(" ");

                // [0] is winning numbers, [1] is attempted numbers
                fillTheSet(s1, s2, arr);

                s1.retainAll(s2);

                if (maxY >= list.size()) {
                    Item it = new Item(maxY + 1, 1);
                    list.add(it);
                }

                // adding
                int size = s1.size();

                Item toAdd = list.get(maxY);
                // System.out.println("I: " + maxY + " size: " + size + " toAdd: " + toAdd.val +
                // " Game: " + toAdd.game);
                // System.out.println("-----");

                for (int j = maxY + 1; j <= size + maxY; j++) {

                    if (j < list.size()) {
                        Item temp = list.get(j);
                        temp.val += toAdd.val;

                    } else {
                        Item temp = new Item(list.get(j - 1).game + 1, 1 + toAdd.val);
                        list.add(temp);
                    }

                }
                /*
                 * 
                 * for (Item integer : list) {
                 * System.out.println(integer.val + " : g -> " + integer.game);
                 * }
                 */
                maxY++;

            }

            for (int i = 0; i < maxY; i++) {
                sum += list.get(i).val;
            }

            System.out.println("Sum: " + sum);

            sc.close();
        } catch (Exception e) {
            System.err.println(e);
        }

    }
}