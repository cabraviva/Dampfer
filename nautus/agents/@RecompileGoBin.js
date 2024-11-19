/*
This file is used to define what happens when you use 'nautus build'
You can write this script like every normal node.js app, but are also
able to use special functions defined below:
async cmd(command: string): Promise<[exitCode, stdout]> - Execute a command in the default shell and waits until completion. Returns [exitCode, stdout]
os(): string - Returns 'windows', 'linux, 'mac' or 'unknown'
info(what: any): void - Displays an info in the console
warn(what: any): void - Displays a warning in the console
error(what: any): void - Displays an error in the console & exits with code 1
exit(code: number = 0): void - Exits with a code (default: 0)
async script(name: string): Promise<void> - Runs another script and returns after it has run. Define it by creating a @ScriptName.js file in this folder and run it by using await script('ScriptName')
async spawn(command: string, args: Array<string>, silent?: bool): Promise<exitCode> - Executes a command and displays the output in the shell
async nodeBin(command: string, args: Array<string>, silent?: bool): Promise<exitCode> - Searches through your locally installed node modules and executes a binary. This can be useful when running pkg, tsc, vite, etc...

modules: A useful collection of some modules, because it's bad practice to use require
Over time we might add more (just check using info(modules)), but right now it's:
- modules.fs
- modules.fse
- modules.path
- modules.chalk
- modules.axios
*/

module.exports = async (cmd, os, info, warn, error, exit, script, spawn, modules, nodeBin) => {

    // This is an agent, define commands to be run if the in
    // agents.yaml specified tank content changes here!

    info('-!-!-!- RECOMPILING GO EXECUTABLE -!-!-!-')

    // Kill every running Dampfer instance
    const killcmd = os() === 'windows' ? `for /f "tokens=2" %i in ('tasklist ^| findstr "Dampfer"') do taskkill /PID %i /F` : `kill -9 $(ps aux | grep '[D]ampfer' | awk '{print $2}')`
    info((await cmd(killcmd))[0] === 0 ? 'Succesfully killed running Dampfer instances' : 'Could not kill Dampfer instances. Might want to check your os or the nautus script!')

    // Build executable
    os() === 'windows' ? await script('BuildGoWinDev') : await script('BuildGoLinuxDev')

    // Run executable
    await script('Run')

}