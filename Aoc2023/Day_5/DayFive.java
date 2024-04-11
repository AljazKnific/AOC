import java.io.File;
import java.util.Arrays;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.Scanner;

class Pair {
    Long start;
    Long end;
    Long diff;

    public Pair(Long start, Long end, Long diff) {
        this.start = start;
        this.end = end;
        this.diff = diff;
    }

    public Pair(Long start, Long end) {
        this.start = start;
        this.end = end;
        this.diff = calcDiff();
    }

    public Long calcDiff() {
        return end - start;
    }

}

class DayFive {

    public static void updateMaps(LinkedList<HashMap<Pair, Pair>> maps, int a, String data) {
        if (maps.size() <= a) {
            maps.add(new HashMap<Pair, Pair>());
        }

        String[] arr = data.split(" ");
        HashMap<Pair, Pair> map = maps.get(a);

        Pair fir = new Pair(Long.parseLong(arr[1]), Long.parseLong(arr[1]) + Long.parseLong(arr[2]),
                Long.parseLong(arr[2]));
        Pair sec = new Pair(Long.parseLong(arr[0]), Long.parseLong(arr[0]) + Long.parseLong(arr[2]),
                Long.parseLong(arr[2]));
        // System.out.println("FIR: " + fir.start + ", " + fir.end + ", " + fir.diff);
        // System.out.println("sec: " + sec.start + ", " + sec.end + ", " + sec.diff);
        map.put(fir, sec);

    }

    public static void main(String[] args) {
        try {
            File inp = new File("Input.txt");
            Scanner sc = new Scanner(inp);

            LinkedList<Long> seeds = new LinkedList<Long>();
            LinkedList<HashMap<Pair, Pair>> maps = new LinkedList<>();
            LinkedList<Pair> seedRanges = new LinkedList<>();
            int type = -1;
            while (sc.hasNextLine()) {
                String data = sc.nextLine();
                // System.out.println(data);

                switch (data) {
                    case "seed-to-soil map:":
                        type = 0;
                        continue;

                    case "soil-to-fertilizer map:":
                        type = 1;
                        continue;

                    case "fertilizer-to-water map:":
                        type = 2;
                        continue;

                    case "water-to-light map:":
                        type = 3;
                        continue;

                    case "light-to-temperature map:":
                        type = 4;
                        continue;

                    case "temperature-to-humidity map:":
                        type = 5;
                        continue;
                    case "humidity-to-location map:":
                        type = 6;
                        continue;

                    case "":
                        type = -2;
                        break;

                    // imamo seeds
                    default:
                        break;
                }

                if (type == -1) {
                    String[] arr = data.split(":");
                    arr = arr[1].split(" ");
                    for (int i = 1; i < arr.length; i++) {
                        seeds.add(Long.parseLong(arr[i]));

                    }

                    for (int i = 1; i < arr.length; i += 2) {
                        seedRanges.add(new Pair(Long.parseLong(arr[i]),
                                Long.parseLong(arr[i]) + Long.parseLong(arr[i + 1]), Long.parseLong(arr[i + 1])));
                    }
                } else if (type > -1) {
                    updateMaps(maps, type, data);
                }

            }
            // PART ONE
            /*
             * 
             * Long bestLoc = Long.MAX_VALUE;
             * for (int j = 0; j < seeds.size(); j++) {
             * Long curr = seeds.get(j);
             * // System.out.println("Seed : " + curr);
             * for (int i = 0; i < maps.size(); i++) {
             * for (Pair p : maps.get(i).keySet()) {
             * if (p.start <= curr && curr < p.end) {
             * Pair two = maps.get(i).get(p);
             * curr = (curr - p.start) + two.start;
             * break;
             * }
             * }
             * 
             * // System.out.println("Curr: " + curr + " Num: " + num);
             * 
             * }
             * 
             * if (curr < bestLoc) {
             * bestLoc = curr;
             * }
             * }
             * 
             * 
             * 
             * System.out.println("Winning seed: " + bestLoc);
             */

            // PART TWO
            LinkedList<Pair> nextRound = new LinkedList<>();

            // cez vse mape
            for (int i = 0; i < maps.size(); i++) {
                // System.out.println("MAPICA: " + i);
                nextRound = new LinkedList<>();
                // cez vse inpute
                while (!seedRanges.isEmpty()) {
                    Pair temp = seedRanges.pop();
                    boolean addOrigi = true;
                    for (Pair p : maps.get(i).keySet()) {

                        // cez vse intervale
                        Pair res = maps.get(i).get(p);

                        // System.out.println("INTERVAL: " + p.start + " " + p.end);
                        // START UNDER END IN
                        if (temp.start < p.start && temp.end >= p.start && temp.end < p.end) {
                            nextRound.add(new Pair(res.start, temp.end - p.start + res.start));
                            seedRanges.add(new Pair(temp.start, p.start - 1));
                            addOrigi = false;
                        } else
                        // BOTH IN
                        if (temp.start >= p.start && temp.end <= p.end) {
                            nextRound.add(new Pair(temp.start - p.start + res.start, temp.end - p.start + res.start));
                            addOrigi = false;

                        } else
                        // START IN END OUT
                        if (temp.start >= p.start && temp.start <= p.end && temp.end > p.end) {
                            nextRound.add(new Pair(temp.start - p.start + res.start, res.end));
                            seedRanges.add(new Pair(p.end + 1, temp.end));
                            addOrigi = false;
                        } else
                        // INTERVAL IS BETWEEN
                        if (temp.start < p.start && temp.end > p.end) {
                            nextRound.add(new Pair(res.start, res.end));
                            seedRanges.add(new Pair(temp.start, p.start - 1));
                            seedRanges.add(new Pair(p.end + 1, temp.end));
                            addOrigi = false;
                        }
                    }

                    if (addOrigi) {
                        nextRound.add(temp);
                    }

                }

                seedRanges = nextRound;

            }

            Long mini = Long.MAX_VALUE;
            for (Pair pair : nextRound) {
                if (pair.start < mini) {
                    mini = pair.start;
                }
            }

            System.out.println("MINI: " + mini);

            sc.close();

        } catch (Exception e) {
            System.err.println(e);
        }

    }
}