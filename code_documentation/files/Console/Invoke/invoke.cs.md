# `Console/Invoke/invoke.cs`

**Functional area:** [17. Thread-safe UI plumbing and shared controls](../../../CODE_OUTLINE.md#17-thread-safe-ui-plumbing-and-shared-controls)

**Role:** Core invoke helpers — run any control update on the UI thread.

## How this file is used

- Used by (incoming references from other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).
- Uses (outgoing references to other files): none found in the graph (file-local or reached via P/Invoke, delegates, or reflection).

## Outline

_Each entry: symbol — line — signature, then a description (from source comments where present, otherwise inferred from naming conventions) and its callers as recorded in the graph._

### Types

#### `UI` (type, L33)

- **`.CallObjectEquals()`** — L46 — `public static object CallObjectEquals(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallObjectGetHashCode()`** — L51 — `public static object CallObjectGetHashCode(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallObjectGetType()`** — L56 — `public static object CallObjectGetType(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallObjectToString()`** — L61 — `public static object CallObjectToString(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallMarshalByRefObjectGetLifetimeService()`** — L77 — `public static object CallMarshalByRefObjectGetLifetimeService(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallMarshalByRefObjectInitializeLifetimeService()`** — L82 — `public static object CallMarshalByRefObjectInitializeLifetimeService(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlAccessibleDefaultActionDescription()`** — L93 — `public static void SetControlAccessibleDefaultActionDescription(Control c, object val)`
  Sets control accessible default action description.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlAccessibleDescription()`** — L98 — `public static void SetControlAccessibleDescription(Control c, object val)`
  Sets control accessible description.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlAccessibleName()`** — L103 — `public static void SetControlAccessibleName(Control c, object val)`
  Sets control accessible name.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlAccessibleRole()`** — L108 — `public static void SetControlAccessibleRole(Control c, object val)`
  Sets control accessible role.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlAllowDrop()`** — L113 — `public static void SetControlAllowDrop(Control c, object val)`
  Sets control allow drop.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlAnchor()`** — L118 — `public static void SetControlAnchor(Control c, object val)`
  Sets control anchor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlBackColor()`** — L123 — `public static void SetControlBackColor(Control c, object val)`
  Sets control back color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlBackgroundImage()`** — L128 — `public static void SetControlBackgroundImage(Control c, object val)`
  Sets control background image.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlBindingContext()`** — L133 — `public static void SetControlBindingContext(Control c, object val)`
  Sets control binding context.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlBounds()`** — L138 — `public static void SetControlBounds(Control c, object val)`
  Sets control bounds.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlCapture()`** — L143 — `public static void SetControlCapture(Control c, object val)`
  Sets control capture.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlCausesValidation()`** — L148 — `public static void SetControlCausesValidation(Control c, object val)`
  Sets control causes validation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlClientSize()`** — L153 — `public static void SetControlClientSize(Control c, object val)`
  Sets control client size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlContextMenuStrip()`** — L158 — `public static void SetControlContextMenuStrip(Control c, object val)`
  Sets control context menu strip.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlCursor()`** — L163 — `public static void SetControlCursor(Control c, object val)`
  Sets control cursor.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlDock()`** — L168 — `public static void SetControlDock(Control c, object val)`
  Sets control dock.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlEnabled()`** — L173 — `public static void SetControlEnabled(Control c, object val)`
  Sets control enabled.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlFont()`** — L178 — `public static void SetControlFont(Control c, object val)`
  Sets control font.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlForeColor()`** — L183 — `public static void SetControlForeColor(Control c, object val)`
  Sets control fore color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlHeight()`** — L188 — `public static void SetControlHeight(Control c, object val)`
  Sets control height.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlImeMode()`** — L193 — `public static void SetControlImeMode(Control c, object val)`
  Sets control ime mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlIsAccessible()`** — L198 — `public static void SetControlIsAccessible(Control c, object val)`
  Sets control is accessible.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlLeft()`** — L203 — `public static void SetControlLeft(Control c, object val)`
  Sets control left.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlLocation()`** — L208 — `public static void SetControlLocation(Control c, object val)`
  Sets control location.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlName()`** — L213 — `public static void SetControlName(Control c, object val)`
  Sets control name.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlParent()`** — L218 — `public static void SetControlParent(Control c, object val)`
  Sets control parent.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlRegion()`** — L223 — `public static void SetControlRegion(Control c, object val)`
  Sets control region.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlRightToLeft()`** — L228 — `public static void SetControlRightToLeft(Control c, object val)`
  Sets control right to left.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlSite()`** — L233 — `public static void SetControlSite(Control c, object val)`
  Sets control site.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlSize()`** — L238 — `public static void SetControlSize(Control c, object val)`
  Sets control size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlTabIndex()`** — L243 — `public static void SetControlTabIndex(Control c, object val)`
  Sets control tab index.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlTabStop()`** — L248 — `public static void SetControlTabStop(Control c, object val)`
  Sets control tab stop.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlTag()`** — L253 — `public static void SetControlTag(Control c, object val)`
  Sets control tag.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlText()`** — L258 — `public static void SetControlText(Control c, object val)`
  Sets control text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlTop()`** — L263 — `public static void SetControlTop(Control c, object val)`
  Sets control top.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlVisible()`** — L268 — `public static void SetControlVisible(Control c, object val)`
  Sets control visible.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetControlWidth()`** — L273 — `public static void SetControlWidth(Control c, object val)`
  Sets control width.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlContains()`** — L282 — `public static object CallControlContains(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlCreateGraphics()`** — L287 — `public static object CallControlCreateGraphics(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlDoDragDrop()`** — L292 — `public static object CallControlDoDragDrop(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlFindForm()`** — L297 — `public static object CallControlFindForm(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlFocus()`** — L302 — `public static object CallControlFocus(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlGetChildAtPoint()`** — L307 — `public static object CallControlGetChildAtPoint(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlGetContainerControl()`** — L312 — `public static object CallControlGetContainerControl(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlGetNextControl()`** — L317 — `public static object CallControlGetNextControl(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlInvalidate()`** — L322 — `public static void CallControlInvalidate(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlPerformLayout()`** — L344 — `public static void CallControlPerformLayout(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlPointToClient()`** — L349 — `public static object CallControlPointToClient(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlPointToScreen()`** — L354 — `public static object CallControlPointToScreen(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlPreProcessMessage()`** — L359 — `public static object CallControlPreProcessMessage(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlRectangleToClient()`** — L365 — `public static object CallControlRectangleToClient(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlRectangleToScreen()`** — L370 — `public static object CallControlRectangleToScreen(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlResumeLayout()`** — L375 — `public static void CallControlResumeLayout(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlScale()`** — L380 — `public static void CallControlScale(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlSelectNextControl()`** — L386 — `public static object CallControlSelectNextControl(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallControlSetBounds()`** — L392 — `public static void CallControlSetBounds(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetScrollableControlAutoScroll()`** — L409 — `public static void SetScrollableControlAutoScroll(Control c, object val)`
  Sets scrollable control auto scroll.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetScrollableControlAutoScrollMargin()`** — L414 — `public static void SetScrollableControlAutoScrollMargin(Control c, object val)`
  Sets scrollable control auto scroll margin.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetScrollableControlAutoScrollMinSize()`** — L419 — `public static void SetScrollableControlAutoScrollMinSize(Control c, object val)`
  Sets scrollable control auto scroll min size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetScrollableControlAutoScrollPosition()`** — L424 — `public static void SetScrollableControlAutoScrollPosition(Control c, object val)`
  Sets scrollable control auto scroll position.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallScrollableControlScrollControlIntoView()`** — L433 — `public static void CallScrollableControlScrollControlIntoView(Control c, object[] val)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallScrollableControlSetAutoScrollMargin()`** — L438 — `public static void CallScrollableControlSetAutoScrollMargin(Control c, object[] val)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetContainerControlActiveControl()`** — L451 — `public static void SetContainerControlActiveControl(Control c, object val)`
  Sets container control active control.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallContainerControlValidate()`** — L460 — `public static object CallContainerControlValidate(Control c, object[] val)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetUpDownBaseBorderStyle()`** — L473 — `public static void SetUpDownBaseBorderStyle(Control c, object val)`
  Sets up down base border style.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetUpDownBaseInterceptArrowKeys()`** — L478 — `public static void SetUpDownBaseInterceptArrowKeys(Control c, object val)`
  Sets up down base intercept arrow keys.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetUpDownBaseReadOnly()`** — L483 — `public static void SetUpDownBaseReadOnly(Control c, object val)`
  Sets up down base read only.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetUpDownBaseTextAlign()`** — L488 — `public static void SetUpDownBaseTextAlign(Control c, object val)`
  Sets up down base text align.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetUpDownBaseUpDownAlign()`** — L493 — `public static void SetUpDownBaseUpDownAlign(Control c, object val)`
  Sets up down base up down align.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallUpDownBaseSelect()`** — L502 — `public static void CallUpDownBaseSelect(Control c, object[] val)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseAcceptsTab()`** — L516 — `public static void SetTextBoxBaseAcceptsTab(Control c, object val)`
  Sets text box base accepts tab.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseAutoSize()`** — L521 — `public static void SetTextBoxBaseAutoSize(Control c, object val)`
  Sets text box base auto size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseBackColor()`** — L526 — `public static void SetTextBoxBaseBackColor(Control c, object val)`
  Sets text box base back color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseBorderStyle()`** — L531 — `public static void SetTextBoxBaseBorderStyle(Control c, object val)`
  Sets text box base border style.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseForeColor()`** — L536 — `public static void SetTextBoxBaseForeColor(Control c, object val)`
  Sets text box base fore color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseHideSelection()`** — L541 — `public static void SetTextBoxBaseHideSelection(Control c, object val)`
  Sets text box base hide selection.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseLines()`** — L546 — `public static void SetTextBoxBaseLines(Control c, object val)`
  Sets text box base lines.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseMaxLength()`** — L551 — `public static void SetTextBoxBaseMaxLength(Control c, object val)`
  Sets text box base max length.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseModified()`** — L556 — `public static void SetTextBoxBaseModified(Control c, object val)`
  Sets text box base modified.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseMultiline()`** — L561 — `public static void SetTextBoxBaseMultiline(Control c, object val)`
  Sets text box base multiline.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseReadOnly()`** — L566 — `public static void SetTextBoxBaseReadOnly(Control c, object val)`
  Sets text box base read only.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseSelectedText()`** — L571 — `public static void SetTextBoxBaseSelectedText(Control c, object val)`
  Sets text box base selected text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseSelectionStart()`** — L576 — `public static void SetTextBoxBaseSelectionStart(Control c, object val)`
  Sets text box base selection start.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxBaseWordWrap()`** — L581 — `public static void SetTextBoxBaseWordWrap(Control c, object val)`
  Sets text box base word wrap.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallTextBoxBaseAppendText()`** — L590 — `public static void CallTextBoxBaseAppendText(Control c, object[] val)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallTextBoxBaseSelect()`** — L595 — `public static void CallTextBoxBaseSelect(Control c, object[] val)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetListControlDataSource()`** — L608 — `public static void SetListControlDataSource(Control c, object val)`
  Sets list control data source.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetListControlDisplayMember()`** — L613 — `public static void SetListControlDisplayMember(Control c, object val)`
  Sets list control display member.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetListControlSelectedValue()`** — L618 — `public static void SetListControlSelectedValue(Control c, object val)`
  Sets list control selected value.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallListControlGetItemText()`** — L627 — `public static object CallListControlGetItemText(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetButtonBaseFlatStyle()`** — L640 — `public static void SetButtonBaseFlatStyle(Control c, object val)`
  Sets button base flat style.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetButtonBaseImage()`** — L645 — `public static void SetButtonBaseImage(Control c, object val)`
  Sets button base image.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetButtonBaseImageAlign()`** — L650 — `public static void SetButtonBaseImageAlign(Control c, object val)`
  Sets button base image align.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetButtonBaseImageIndex()`** — L655 — `public static void SetButtonBaseImageIndex(Control c, object val)`
  Sets button base image index.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetButtonBaseImageList()`** — L660 — `public static void SetButtonBaseImageList(Control c, object val)`
  Sets button base image list.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetButtonBaseImeMode()`** — L665 — `public static void SetButtonBaseImeMode(Control c, object val)`
  Sets button base ime mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetButtonBaseTextAlign()`** — L670 — `public static void SetButtonBaseTextAlign(Control c, object val)`
  Sets button base text align.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetButtonDialogResult()`** — L683 — `public static void SetButtonDialogResult(Control c, object val)`
  Sets button dialog result.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallButtonNotifyDefault()`** — L692 — `public static void CallButtonNotifyDefault(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxBackColor()`** — L705 — `public static void SetComboBoxBackColor(Control c, object val)`
  Sets combo box back color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxDrawMode()`** — L710 — `public static void SetComboBoxDrawMode(Control c, object val)`
  Sets combo box draw mode.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxDropDownStyle()`** — L715 — `public static void SetComboBoxDropDownStyle(Control c, object val)`
  Sets combo box drop down style.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxDropDownWidth()`** — L720 — `public static void SetComboBoxDropDownWidth(Control c, object val)`
  Sets combo box drop down width.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxDroppedDown()`** — L725 — `public static void SetComboBoxDroppedDown(Control c, object val)`
  Sets combo box dropped down.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxForeColor()`** — L730 — `public static void SetComboBoxForeColor(Control c, object val)`
  Sets combo box fore color.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxIntegralHeight()`** — L735 — `public static void SetComboBoxIntegralHeight(Control c, object val)`
  Sets combo box integral height.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxMaxDropDownItems()`** — L740 — `public static void SetComboBoxMaxDropDownItems(Control c, object val)`
  Sets combo box max drop down items.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxMaxLength()`** — L745 — `public static void SetComboBoxMaxLength(Control c, object val)`
  Sets combo box max length.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxSelectedIndex()`** — L750 — `public static void SetComboBoxSelectedIndex(Control c, object val)`
  Sets combo box selected index.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxSelectedItem()`** — L755 — `public static void SetComboBoxSelectedItem(Control c, object val)`
  Sets combo box selected item.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxSelectedText()`** — L760 — `public static void SetComboBoxSelectedText(Control c, object val)`
  Sets combo box selected text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxSelectionLength()`** — L765 — `public static void SetComboBoxSelectionLength(Control c, object val)`
  Sets combo box selection length.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxSelectionStart()`** — L770 — `public static void SetComboBoxSelectionStart(Control c, object val)`
  Sets combo box selection start.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxSorted()`** — L775 — `public static void SetComboBoxSorted(Control c, object val)`
  Sets combo box sorted.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxText()`** — L780 — `public static void SetComboBoxText(Control c, object val)`
  Sets combo box text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetComboBoxValueMember()`** — L785 — `public static void SetComboBoxValueMember(Control c, object val)`
  Sets combo box value member.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallComboBoxFindString()`** — L794 — `public static object CallComboBoxFindString(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallComboBoxFindStringExact()`** — L802 — `public static object CallComboBoxFindStringExact(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallComboBoxGetItemHeight()`** — L810 — `public static object CallComboBoxGetItemHeight(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxAcceptsReturn()`** — L823 — `public static void SetTextBoxAcceptsReturn(Control c, object val)`
  Sets text box accepts return.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxCharacterCasing()`** — L828 — `public static void SetTextBoxCharacterCasing(Control c, object val)`
  Sets text box character casing.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxPasswordChar()`** — L833 — `public static void SetTextBoxPasswordChar(Control c, object val)`
  Sets text box password char.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxScrollBars()`** — L838 — `public static void SetTextBoxScrollBars(Control c, object val)`
  Sets text box scroll bars.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxSelectionLength()`** — L843 — `public static void SetTextBoxSelectionLength(Control c, object val)`
  Sets text box selection length.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxText()`** — L848 — `public static void SetTextBoxText(Control c, object val)`
  Sets text box text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTextBoxTextAlign()`** — L853 — `public static void SetTextBoxTextAlign(Control c, object val)`
  Sets text box text align.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCheckBoxAppearance()`** — L866 — `public static void SetCheckBoxAppearance(Control c, object val)`
  Sets check box appearance.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCheckBoxAutoCheck()`** — L871 — `public static void SetCheckBoxAutoCheck(Control c, object val)`
  Sets check box auto check.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCheckBoxCheckAlign()`** — L876 — `public static void SetCheckBoxCheckAlign(Control c, object val)`
  Sets check box check align.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCheckBoxChecked()`** — L881 — `public static void SetCheckBoxChecked(Control c, object val)`
  Sets check box checked.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCheckBoxCheckState()`** — L886 — `public static void SetCheckBoxCheckState(Control c, object val)`
  Sets check box check state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetCheckBoxThreeState()`** — L891 — `public static void SetCheckBoxThreeState(Control c, object val)`
  Sets check box three state.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRadioButtonAppearance()`** — L904 — `public static void SetRadioButtonAppearance(Control c, object val)`
  Sets radio button appearance.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRadioButtonAutoCheck()`** — L909 — `public static void SetRadioButtonAutoCheck(Control c, object val)`
  Sets radio button auto check.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRadioButtonCheckAlign()`** — L914 — `public static void SetRadioButtonCheckAlign(Control c, object val)`
  Sets radio button check align.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRadioButtonChecked()`** — L919 — `public static void SetRadioButtonChecked(Control c, object val)`
  Sets radio button checked.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetRadioButtonTextAlign()`** — L924 — `public static void SetRadioButtonTextAlign(Control c, object val)`
  Sets radio button text align.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetNumericUpDownDecimalPlaces()`** — L937 — `public static void SetNumericUpDownDecimalPlaces(Control c, object val)`
  Sets numeric up down decimal places.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetNumericUpDownHexadecimal()`** — L942 — `public static void SetNumericUpDownHexadecimal(Control c, object val)`
  Sets numeric up down hexadecimal.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetNumericUpDownIncrement()`** — L947 — `public static void SetNumericUpDownIncrement(Control c, object val)`
  Sets numeric up down increment.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetNumericUpDownMaximum()`** — L952 — `public static void SetNumericUpDownMaximum(Control c, object val)`
  Sets numeric up down maximum.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetNumericUpDownMinimum()`** — L957 — `public static void SetNumericUpDownMinimum(Control c, object val)`
  Sets numeric up down minimum.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetNumericUpDownText()`** — L962 — `public static void SetNumericUpDownText(Control c, object val)`
  Sets numeric up down text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetNumericUpDownThousandsSeparator()`** — L967 — `public static void SetNumericUpDownThousandsSeparator(Control c, object val)`
  Sets numeric up down thousands separator.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetNumericUpDownValue()`** — L972 — `public static void SetNumericUpDownValue(Control c, object val)`
  Sets numeric up down value.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTrackBarAutoSize()`** — L985 — `public static void SetTrackBarAutoSize(Control c, object val)`
  Sets track bar auto size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTrackBarLargeChange()`** — L990 — `public static void SetTrackBarLargeChange(Control c, object val)`
  Sets track bar large change.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTrackBarMaximum()`** — L995 — `public static void SetTrackBarMaximum(Control c, object val)`
  Sets track bar maximum.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTrackBarMinimum()`** — L1000 — `public static void SetTrackBarMinimum(Control c, object val)`
  Sets track bar minimum.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTrackBarOrientation()`** — L1005 — `public static void SetTrackBarOrientation(Control c, object val)`
  Sets track bar orientation.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTrackBarSmallChange()`** — L1010 — `public static void SetTrackBarSmallChange(Control c, object val)`
  Sets track bar small change.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTrackBarTickFrequency()`** — L1015 — `public static void SetTrackBarTickFrequency(Control c, object val)`
  Sets track bar tick frequency.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTrackBarTickStyle()`** — L1020 — `public static void SetTrackBarTickStyle(Control c, object val)`
  Sets track bar tick style.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetTrackBarValue()`** — L1025 — `public static void SetTrackBarValue(Control c, object val)`
  Sets track bar value.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.CallTrackBarSetRange()`** — L1034 — `public static void CallTrackBarSetRange(Control c, object[] obj)`
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLabelAutoSize()`** — L1047 — `public static void SetLabelAutoSize(Control c, object val)`
  Sets label auto size.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLabelBorderStyle()`** — L1052 — `public static void SetLabelBorderStyle(Control c, object val)`
  Sets label border style.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLabelFlatStyle()`** — L1057 — `public static void SetLabelFlatStyle(Control c, object val)`
  Sets label flat style.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLabelImage()`** — L1062 — `public static void SetLabelImage(Control c, object val)`
  Sets label image.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLabelImageAlign()`** — L1067 — `public static void SetLabelImageAlign(Control c, object val)`
  Sets label image align.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLabelImageIndex()`** — L1072 — `public static void SetLabelImageIndex(Control c, object val)`
  Sets label image index.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLabelImageList()`** — L1077 — `public static void SetLabelImageList(Control c, object val)`
  Sets label image list.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLabelText()`** — L1082 — `public static void SetLabelText(Control c, object val)`
  Sets label text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLabelTextAlign()`** — L1087 — `public static void SetLabelTextAlign(Control c, object val)`
  Sets label text align.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetLabelUseMnemonic()`** — L1092 — `public static void SetLabelUseMnemonic(Control c, object val)`
  Sets label use mnemonic.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetGroupBoxAllowDrop()`** — L1105 — `public static void SetGroupBoxAllowDrop(Control c, object val)`
  Sets group box allow drop.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetGroupBoxFlatStyle()`** — L1110 — `public static void SetGroupBoxFlatStyle(Control c, object val)`
  Sets group box flat style.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetGroupBoxText()`** — L1115 — `public static void SetGroupBoxText(Control c, object val)`
  Sets group box text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPanelBorderStyle()`** — L1128 — `public static void SetPanelBorderStyle(Control c, object val)`
  Sets panel border style.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPanelTabStop()`** — L1133 — `public static void SetPanelTabStop(Control c, object val)`
  Sets panel tab stop.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.
- **`.SetPanelText()`** — L1138 — `public static void SetPanelText(Control c, object val)`
  Sets panel text.
  No callers found in the graph — likely invoked via P/Invoke, UI/event wiring, a delegate, a thread start, or externally.

---
_Generated from the graphify knowledge graph (`graphify-out/graph.json`); line numbers refer to `Project Files/Source/Console/Invoke/invoke.cs`. Regenerate after code changes with `graphify update "Project Files/Source"` followed by `python code_documentation/tools/gen_file_docs.py`._
