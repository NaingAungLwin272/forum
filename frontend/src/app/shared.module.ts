import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { NgZorroAntdModule } from './ng-zorro-antd/ng-zorro-antd.module';
import { QuestionTableComponent } from './components/question-table/question-table.component';
import { ScrollingModule } from '@angular/cdk/scrolling';
import { ListTableComponent } from './components/list-table/list-table.component';
import { CKEditorModule } from '@ckeditor/ckeditor5-angular';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommentComponent } from './components/comment/comment.component';
import { CkeEditorComponent } from './components/cke-editor/cke-editor.component';
import { CommentDetailComponent } from './components/comment-detail/comment-detail.component';
import { TranslateModule } from '@ngx-translate/core';
import { ChangePasswordComponent } from './components/change-password/change-password.component';


@NgModule({
  declarations: [
    QuestionTableComponent,
    ListTableComponent,
    CommentComponent,
    CkeEditorComponent,
    CommentDetailComponent,
    ChangePasswordComponent
  ],
  imports: [
    NgZorroAntdModule,
    RouterModule,
    ScrollingModule,
    CKEditorModule,
    FormsModule,
    ReactiveFormsModule,
    TranslateModule,

  ],
  exports: [
    NgZorroAntdModule,
    QuestionTableComponent,
    ListTableComponent,
    CKEditorModule,
    FormsModule,
    ReactiveFormsModule,
    CommentComponent,
    CkeEditorComponent,
    ChangePasswordComponent,
    TranslateModule
  ]
})
export class SharedModule { }
