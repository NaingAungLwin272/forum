import { Component, OnInit, inject } from '@angular/core';
import * as customBuild from '../../components/ckCus/build/ckeditor';
import { ThemeService } from 'src/app/services/theme.service';
import { UserModalData } from 'src/app/interfaces/user';
import { NZ_MODAL_DATA, NzModalRef } from 'ng-zorro-antd/modal';

@Component({
  selector: 'app-comment-detail',
  templateUrl: './comment-detail.component.html',
  styleUrls: ['./comment-detail.component.scss']
})
export class CommentDetailComponent implements OnInit {
  darkTheme!: boolean;
  readonly userModalData: UserModalData = inject(NZ_MODAL_DATA);
  editor = customBuild;
  editorConfig = {
    toolbar: {
      items: [],
      shouldNotGroupWhenFull: true,
    },
  };
  constructor(
    private themeService: ThemeService,
    private modalRef: NzModalRef,
  ) { }

  ngOnInit(): void {

    this.darkTheme = localStorage.getItem('theme') === 'dark' ? true : false;
    enum ThemeType {
      dark = 'dark',
      default = 'default',
    }
    this.themeService.currentTheme$.subscribe((currentTheme: ThemeType) => {
      this.darkTheme = currentTheme === ThemeType.dark;
    });
  }

  cancel(): void {
    this.modalRef.destroy("cancel");
  }
}
