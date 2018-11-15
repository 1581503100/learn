package lsj.cc.flag;

/**
 * @Author:sjliu7
 * @Date:2018/11/14 16:10
 */
public class BoolFlag {
    String name;
    boolean def;
    String usage;

    public BoolFlag(String name, boolean def, String usage) {
        this.name = name;
        this.def = def;
        this.usage = usage;
    }

    public String getName() {
        return name;
    }

    public BoolFlag name(String name) {
        this.name = name;
        return this;
    }

    public boolean isDef() {
        return def;
    }

    public BoolFlag def(boolean def) {
        this.def = def;
        return this;
    }

    public String getUsage() {
        return usage;
    }

    public BoolFlag usage(String usage) {
        this.usage = usage;
        return this;
    }
}
