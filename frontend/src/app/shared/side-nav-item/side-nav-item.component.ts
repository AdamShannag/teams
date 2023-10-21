import { Component, Input } from '@angular/core';

@Component({
  selector: 'side-nav-item',
  templateUrl: './side-nav-item.component.html',
  styleUrls: ['./side-nav-item.component.scss'],
})
export class SideNavItemComponent {
  @Input({ required: true })
  link!: string;
  @Input({ required: true })
  icon!: string;
  @Input({ required: true })
  title!: string;
  @Input()
  show: boolean | null = true;
}
