//Problem
void renderBanner() {
    if ((platform.toUpperCase().indexOf("MAC") > -1) &&
         (browser.toUpperCase().indexOf("IE") > -1) &&
          wasInitialized() && resize > 0 )
    {
      // do something
    }
  }

//Solution
void renderBanner() {
  final boolean isMacOs = platform.toUpperCase().indexOf("MAC") > -1;
  final boolean isIE = browser.toUpperCase().indexOf("IE") > -1;
  final boolean wasResized = resize > 0;

  if (isMacOs && isIE && wasInitialized() && wasResized) {
    // do something
  }
}