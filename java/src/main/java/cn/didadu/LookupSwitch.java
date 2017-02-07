package cn.didadu;

/**
 * Created by zhangjing on 17-2-7.
 */
public class LookupSwitch {

    int chooseFar(int i){
        switch (i) {
            case -100: return -1;
            case 0: return 0;
            case 100: return 1;
            default: return -1;
        }
    }

}
