import { Directive, HostListener } from '@angular/core';

@Directive({
  selector: '[appCkeNoFocus]'
})
export class CkeNoFocusDirective {
  @HostListener('mousedown', ['$event'])
  onMouseDown(event: MouseEvent) {
    event.preventDefault();
  }
}
