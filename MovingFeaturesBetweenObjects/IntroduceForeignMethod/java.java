//Problem
class Report {
    // ...
    void sendReport() {
      Date nextDay = new Date(previousEnd.getYear(),
        previousEnd.getMonth(), previousEnd.getDate() + 1);
      // ...
    }
  }


//Solution
class Report {
    // ...
    void sendReport() {
      Date newStart = nextDay(previousEnd);
      // ...
    }
    private static Date nextDay(Date arg) {
      return new Date(arg.getYear(), arg.getMonth(), arg.getDate() + 1);
    }
  }