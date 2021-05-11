# procrastination-killer

This app finds vscode process code and check wheter is running or not.

If it's not running and X time passes, then it sends a mail by using gooogle smtp to notify that the program was not opened in X hours
and in consequence i owe an "item".

If it finds vscode process, but it is closed before X time, then the punish email will be triggered.
