package cn.didadu;

/**
 * Created by zhangjing on 17-2-6.
 */
public class TableSwitch {

    int chooseNear(int i){
        switch (i) {
            case 0: return 0;
            case 1: return 1;
            case 2: return 2;
            default: return -1;
        }
    }

}
