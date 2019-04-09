# vscode keys

- box selection = shift + option + mouse
- edit multiple lines = option + command + up/down arrow
- trigger left panel = command + b
- disable explorer window to reveal file for current opening file = uncheck autoReveal
- manual suggestion trigger = has beed set to cmd + 1 in keyboard shortcuts
- fix auto complete issue

        Run Go: Install/Update Tools in VSCode
        cd ~/go/bin
        ./gocode close


## Basic Editing
Key	Command	Command id
⌘X	Cut line (empty selection)	editor.action.clipboardCutAction
⌘C	Copy line (empty selection)	editor.action.clipboardCopyAction
⇧⌘K	Delete Line	editor.action.deleteLines
⌘Enter	Insert Line Below	editor.action.insertLineAfter
⇧⌘Enter	Insert Line Above	editor.action.insertLineBefore
⌥↓	Move Line Down	editor.action.moveLinesDownAction
⌥↑	Move Line Up	editor.action.moveLinesUpAction
⇧⌥↓	Copy Line Down	editor.action.copyLinesDownAction
⇧⌥↑	Copy Line Up	editor.action.copyLinesUpAction
⌘D	Add Selection To Next Find Match	editor.action.addSelectionToNextFindMatch
⌘K ⌘D	Move Last Selection To Next Find Match	editor.action.moveSelectionToNextFindMatch
⌘U	Undo last cursor operation	cursorUndo
⇧⌥I	Insert cursor at end of each line selected	editor.action.insertCursorAtEndOfEachLineSelected
⇧⌘L	Select all occurrences of current selection	editor.action.selectHighlights
⌘F2	Select all occurrences of current word	editor.action.changeAll
⌘L	Select current line	expandLineSelection
⌥⌘↓	Insert Cursor Below	editor.action.insertCursorBelow
⌥⌘↑	Insert Cursor Above	editor.action.insertCursorAbove
⇧⌘\	Jump to matching bracket	editor.action.jumpToBracket
⌘]	Indent Line	editor.action.indentLines
⌘[	Outdent Line	editor.action.outdentLines
Home	Go to Beginning of Line	cursorHome
End	Go to End of Line	cursorEnd
⌘↓	Go to End of File	cursorBottom
⌘↑	Go to Beginning of File	cursorTop
⌃PageDown	Scroll Line Down	scrollLineDown
⌃PageUp	Scroll Line Up	scrollLineUp
⌘PageDown	Scroll Page Down	scrollPageDown
⌘PageUp	Scroll Page Up	scrollPageUp
⌥⌘[	Fold (collapse) region	editor.fold
⌥⌘]	Unfold (uncollapse) region	editor.unfold
⌘K ⌘[	Fold (collapse) all subregions	editor.foldRecursively
⌘K ⌘]	Unfold (uncollapse) all subregions	editor.unfoldRecursively
⌘K ⌘0	Fold (collapse) all regions	editor.foldAll
⌘K ⌘J	Unfold (uncollapse) all regions	editor.unfoldAll
⌘K ⌘C	Add Line Comment	editor.action.addCommentLine
⌘K ⌘U	Remove Line Comment	editor.action.removeCommentLine
⌘/	Toggle Line Comment	editor.action.commentLine
⇧⌥A	Toggle Block Comment	editor.action.blockComment
⌘F	Find	actions.find
⌥⌘F	Replace	editor.action.startFindReplaceAction
⌘G	Find Next	editor.action.nextMatchFindAction
⇧⌘G	Find Previous	editor.action.previousMatchFindAction
⌥Enter	Select All Occurrences of Find Match	editor.action.selectAllMatches
⌥⌘C	Toggle Find Case Sensitive	toggleFindCaseSensitive
⌥⌘R	Toggle Find Regex	toggleFindRegex
⌥⌘W	Toggle Find Whole Word	toggleFindWholeWord
⌃⇧M	Toggle Use of Tab Key for Setting Focus	editor.action.toggleTabFocusMode
unassigned	Toggle Render Whitespace	toggleRenderWhitespace
⌥Z	Toggle Word Wrap	editor.action.toggleWordWrap


## Rich Languages Editing
Key	Command	Command id
⌃Space	Trigger Suggest	editor.action.triggerSuggest
⇧⌘Space	Trigger Parameter Hints	editor.action.triggerParameterHints
⇧⌥F	Format Document	editor.action.formatDocument
⌘K ⌘F	Format Selection	editor.action.formatSelection
F12	Go to Definition	editor.action.revealDefinition
⌘K ⌘I	Show Hover	editor.action.showHover
⌥F12	Peek Definition	editor.action.peekDefinition
⌘K F12	Open Definition to the Side	editor.action.revealDefinitionAside
⌘.	Quick Fix	editor.action.quickFix
⇧F12	Peek References	editor.action.referenceSearch.trigger
F2	Rename Symbol	editor.action.rename
⇧⌘.	Replace with Next Value	editor.action.inPlaceReplace.down
⇧⌘,	Replace with Previous Value	editor.action.inPlaceReplace.up
⌃⇧⌘→	Expand AST Selection	editor.action.smartSelect.expand
⌃⇧⌘←	Shrink AST Selection	editor.action.smartSelect.shrink
⌘K ⌘X	Trim Trailing Whitespace	editor.action.trimTrailingWhitespace
⌘K M	Change Language Mode	workbench.action.editor.changeLanguageMode
Navigation
Key	Command	Command id
⌘T	Show All Symbols	workbench.action.showAllSymbols
⌃G	Go to Line...	workbench.action.gotoLine
⌘P	Go to File..., Quick Open	workbench.action.quickOpen
⇧⌘O	Go to Symbol...	workbench.action.gotoSymbol
⇧⌘M	Show Problems	workbench.actions.view.problems
F8	Go to Next Error or Warning	editor.action.marker.nextInFiles
⇧F8	Go to Previous Error or Warning	editor.action.marker.prevInFiles
⇧⌘P	Show All Commands	workbench.action.showCommands
⌃⇧Tab	Navigate Editor Group History	workbench.action.openPreviousRecentlyUsedEditorInGroup
⌃-	Go Back	workbench.action.navigateBack
⌃-	Go back in Quick Input	workbench.action.quickInputBack
⌃⇧-	Go Forward	workbench.action.navigateForward
