package lsj.cc;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Fan {

    public static enum A{
        A,B,C,D
    }
    public static void main(String[] args) {
        Map<A,String >  map = new HashMap<>();
        map.put(A.A,"A");
        map.put(A.B,"B");
        System.out.println(map.get(A.B));

    }



}
