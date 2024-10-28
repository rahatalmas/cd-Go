class Employee{
    private String name;
    private String designation;
    private int totalLeave;
    private int sickLeave;
    private int casualLeave;
    private int usedSickLeave;
    private int usedCasualLeave;

    Employee(String name,String designation){
        this.name = name;
        this.designation = designation;
        this.totalLeave = 35;
        this.sickLeave = 15;
        this.casualLeave = 20;
        this.usedSickLeave = 0;
        this.usedCasualLeave = 0;
    }
    public String getName(){
        return name;
    }
    public String getDesignation(){
        return designation;
    }
    public int getTotalLeave(){
        return totalLeave;
    }
    public int getSickLeave(){
        return sickLeave;
    }
    public int getCasualLeave(){
        return casualLeave;
    }
    public int getUsedSickLeave(){
        return usedSickLeave;
    }
    public int getUsedCasualLeave(){
        return usedCasualLeave;
    }
    public void setUsedSickLeave(int days){
        this.usedSickLeave = days;
    }
    public void setUsedCasualLeave(int days){
        this.usedCasualLeave = days;
    }
    public void printInfo(){
        System.out.println("Name: "+name);
        System.out.println("Designation: "+designation);
        System.out.println("Total Leave: "+totalLeave);
        System.out.println("Sick Leave: "+sickLeave);
        System.out.println("Casual Leave: "+casualLeave);
        System.out.println("Used Sick Leave: "+usedSickLeave);
        System.out.println("Used Casual Leave: "+usedCasualLeave);
    }
}

class Executive extends Employee{
       private int maxCasualLeaveAtOnce;
       Executive(String name,String designation){
         super(name,designation);
         this.maxCasualLeaveAtOnce = 10;
       }
       public void requestLeave(String leaveType,int days){
          int totalleave = getTotalLeave();
          int casualleave = getCasualLeave();
          int sickleave = getSickLeave();

          System.out.println("Dear "+getName()+", ");
          System.out.println("Designation: "+getDesignation());
          if(leaveType == "casual"){
             if(days > maxCasualLeaveAtOnce){
                System.out.println("Sorry, you can take max 10 days casual leave at once...");
                System.out.println("Your remain casual leave: "+(getCasualLeave()-getUsedCasualLeave())+" days...");
                System.out.println("Your remain sick leave: "+(getSickLeave()-getUsedSickLeave())+" days...\n");
                return;
             }else if((getUsedCasualLeave() + days)<=casualleave){
                setUsedCasualLeave(getUsedCasualLeave()+days);
             }else{
                System.out.println("Sorry, You don't have "+days+" days casual leave remain...");
                System.out.println("Your remain casual leave: "+(getCasualLeave()-getUsedCasualLeave())+" days...");
                System.out.println("Your remain sick leave: "+(getSickLeave()-getUsedSickLeave())+" days...\n");
                return;
             }
          }else{
             int remaintotalleave = totalleave-(getUsedCasualLeave()+getUsedSickLeave());
             if(sickleave>=(getUsedSickLeave()+days)){
                setUsedSickLeave(getUsedSickLeave()+days);
             }else if(remaintotalleave>=days){
                setUsedSickLeave(sickleave);
                setUsedCasualLeave(getUsedCasualLeave()+(days-sickleave));
             }else{
                System.out.println("Sorry, You can't take "+days+" leave...");
                System.out.println("Your remain casual leave: "+(getCasualLeave()-getUsedCasualLeave())+" days...");
                System.out.println("Your remain sick leave: "+(getSickLeave()-getUsedSickLeave())+" days...\n");
                return;
             }
          }
          System.out.println(days+" days "+leaveType+" leave granted...");
          System.out.println("Your remain casual leave: "+(getCasualLeave()-getUsedCasualLeave())+" days...");
          System.out.println("Your remain sick leave: "+(getSickLeave()-getUsedSickLeave())+" days...");;
          System.out.print("\n");
       }
       public void cancelLeave(String leaveType,int days){
          if(leaveType == "casual"){
             if(days>getUsedCasualLeave()){
                System.out.println("Invalid Request, you didn't take "+days+" Casual leave");
                return;
             }else{
                setUsedCasualLeave(getUsedCasualLeave()-days);
             }
          }else{
             int all = getUsedSickLeave() + getUsedCasualLeave();
             if(days>all){
                System.out.println("Invalid Request, you didn't take "+days+" sick leave");
                return;
             }else if(days<=getUsedSickLeave()){
                setUsedSickLeave(getUsedSickLeave()-days);
             }else{
                setUsedSickLeave(0);
                setUsedCasualLeave(getUsedCasualLeave()-(days-getSickLeave()));
             }
          }
          System.out.println("Your "+days+" days leave cancel request accepted");
          System.out.println("Remaining leave: ");
          System.out.println("Casual leave: "+(getCasualLeave()-getUsedCasualLeave()));
          System.out.println("Sick leave: "+(getSickLeave()-getUsedSickLeave()));
          System.out.print("\n");
       }
       public void printInfo(){
          super.printInfo();
          System.out.println("Max Casual Leave At Once: "+maxCasualLeaveAtOnce+"\n");
       }
}

class Manager extends Employee{
       private int maxCasualLeaveAtOnce;
       Manager(String name,String designation){
         super(name,designation);
         this.maxCasualLeaveAtOnce = 15;
       }
       public void requestLeave(String leaveType,int days){
          int totalleave = getTotalLeave();
          int casualleave = getCasualLeave();
          int sickleave = getSickLeave();

          System.out.println("Dear "+getName()+", ");
          System.out.println("Designation: "+getDesignation());
          if(leaveType == "casual"){
             if(days > maxCasualLeaveAtOnce){
                System.out.println("Sorry, you can take max 10 days casual leave at once...");
                System.out.println("Your remain casual leave: "+(getCasualLeave()-getUsedCasualLeave())+" days...");
                System.out.println("Your remain sick leave: "+(getSickLeave()-getUsedSickLeave())+" days...\n");
                return;
             }else if((getUsedCasualLeave() + days)<=casualleave){
                setUsedCasualLeave(getUsedCasualLeave()+days);
             }else{
                System.out.println("Sorry, You don't have "+days+" days casual leave remain...");
                System.out.println("Your remain casual leave: "+(getCasualLeave()-getUsedCasualLeave())+" days...");
                System.out.println("Your remain sick leave: "+(getSickLeave()-getUsedSickLeave())+" days...\n");
                return;
             }
          }else{
             int remaintotalleave = totalleave-(getUsedCasualLeave()+getUsedSickLeave());
             if(sickleave>=(getUsedSickLeave()+days)){
                setUsedSickLeave(getUsedSickLeave()+days);
             }else if(remaintotalleave>=days){
                setUsedSickLeave(sickleave);
                setUsedCasualLeave(getUsedCasualLeave()+(days-sickleave));
             }else{
                System.out.println("Sorry, You can't take "+days+" leave...");
                System.out.println("Your remain casual leave: "+(getCasualLeave()-getUsedCasualLeave())+" days...");
                System.out.println("Your remain sick leave: "+(getSickLeave()-getUsedSickLeave())+" days...\n");
                return;
             }
          }
          System.out.println(days+" days "+leaveType+" leave granted...");
          System.out.println("Your remain casual leave: "+(getCasualLeave()-getUsedCasualLeave())+" days...");
          System.out.println("Your remain sick leave: "+(getSickLeave()-getUsedSickLeave())+" days...");;
          System.out.print("\n");
       }
       public void printInfo(){
          super.printInfo();
          System.out.println("Max Casual Leave At Once: "+maxCasualLeaveAtOnce+"\n");
       }
}

public class main {
    public static void main(String[] args) {
        Executive executive = new Executive("Alice","Executive");
        Manager manager = new Manager("Bob","Manager");
        executive.requestLeave("sick",5);
        executive.requestLeave("casual",8);
        executive.requestLeave("casual",11);
        executive.requestLeave("sick",18);
        
        manager.requestLeave("sick",3);
        manager.requestLeave("casual",15);
        manager.requestLeave("casual",16);

        executive.printInfo();
        manager.printInfo();

        executive.cancelLeave("sick",2);
        executive.cancelLeave("casual",3);
        executive.cancelLeave("casual",6);

        executive.printInfo();
    }
}
