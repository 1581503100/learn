package lsj.cc.flag;

import java.util.HashMap;
import java.util.Map;

/**
 * @Author:sjliu7
 * @Date:2018/11/14 15:47
 */
public class Flag {

    private static Map<String,Object>  map = new HashMap<>();


    public static void parser(StringFlag[] flags,String ...args){
        for(StringFlag flag:flags){
            map.put(flag.getName(),flag.getUsage());
        }

        for (int i=0;i<args.length;i++){
            String a  = args[i];
            if(a.startsWith("-")){
                a =trim(a);
                if(a.contains("=")){
                    String []kv = a.split(a,2);
                    Object igl = map.get(kv[0]);
                    if(igl == null || kv[0].equals("")){
                        System.exit(1);
                    }
                    map.put(kv[0],kv[1]);
                }
                else{
                    if(i+1<args.length){

                    }
                }
            }
        }
    }

    private static String trim(String  s){
        int idx= 0;
        for (int i=0;i<s.length();i++){
            if(s.charAt(i) == '-'){
                idx++;
            }
            else{
                break;
            }
        }
        return s.substring(idx);
    }
}
