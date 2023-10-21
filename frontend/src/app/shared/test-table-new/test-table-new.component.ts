import { AfterViewInit, Component, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { SelectionModel } from '@angular/cdk/collections';
import { ActivatedRoute } from '@angular/router';
import { tap, merge } from 'rxjs';

@Component({
  selector: 'test-table-new',
  templateUrl: './test-table-new.component.html',
  styleUrls: ['./test-table-new.component.scss'],
})
export class TestTableNewComponent implements AfterViewInit {
  @ViewChild(MatPaginator)
  paginator!: MatPaginator;
  @ViewChild(MatSort)
  sort!: MatSort;

  headers: { key: string; label: string }[] = [];
  data: any[] = [];
  dataSize = 0;

  displayedColumns = ['select'];
  selection = new SelectionModel<any>(true, []);

  constructor(private route: ActivatedRoute) {
    this.headers = this.route.snapshot.data['tableData'].headers;
    this.data = this.route.snapshot.data['tableData'].data;

    this.headers.map((item) => this.displayedColumns.push(item.key));
  }

  onDataToggled(data: any) {
    this.selection.toggle(data);
  }

  isAllSelected() {
    return this.selection.selected?.length == this.data?.length;
  }

  toggleAll() {
    if (this.isAllSelected()) {
      this.selection.clear();
    } else {
      this.selection.select(...this.data);
    }
  }

  loadPageData() {
  }

  ngAfterViewInit() {
    this.sort.sortChange.subscribe(() => {
      this.paginator.pageIndex = 0;
    });
    merge(this.sort.sortChange, this.paginator.page)
      .pipe(tap(() => this.loadPageData()))
      .subscribe();
  }
}
