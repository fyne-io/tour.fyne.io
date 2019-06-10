---
layout: home

title: Application and RunLoop
order: 1

redirect_from:
- /basics/

---

For a GUI application to work it needs to run an event loop
(sometimes called a runloop) that processes user interactions
and drawing events. In Fyne this is started using the `App.Run()`
or `Window.ShowAndRun()` functions. One of these must be called
from the end of your setup code in the `main()` function.

An application should only have one runloop and so you should only
call `Run()` once in your code. Calling it a second time will cause
errors.

The helper method `ShowAndRun()` on `fyne.Window` allows you to
show your window and run the application at the same time.
If you wish to show a second window you must only call the `Show()`
function. This is illustrated in the `showAnother` function.

See also that functions executed after `Run()` will not be called
until the application exits.