import { Component, Input, OnInit } from '@angular/core';
import * as customBuild from '../ckCus/build/ckeditor';
import { User } from 'src/app/interfaces/user';
import { EditorConfig } from 'src/app/interfaces/mention';
import { ControlContainer, FormGroup } from '@angular/forms';
import { BehaviorSubject } from 'rxjs';


@Component({
  selector: 'app-cke-editor',
  templateUrl: './cke-editor.component.html',
  styleUrls: ['./cke-editor.component.scss']
})
export class CkeEditorComponent implements OnInit {
  excludeUser!: User[];
  private _userData = new BehaviorSubject<User[]>([]);
  @Input() showToolbar = true;
  @Input() disabled = true;
  @Input()
  set userData(value) { this._userData.next(value); }
  get userData() { return this._userData.getValue(); }
  public editorGroup!: FormGroup;
  public editor = customBuild;
  darkTheme!: boolean;
  cb = function () { return (new Date()).getTime(); }

  constructor(
    private controlContainer: ControlContainer,
  ) { }

  ngOnInit(): void {
    this.editorGroup = <FormGroup>this.controlContainer.control;
    const currentUser = localStorage.getItem('login_user_id');
    this._userData.subscribe(() => {
      this.excludeUser = this.userData.filter(user => user._id != currentUser)
      this.editorConfig.mention.feeds = [
        { marker: '@', feed: this.excludeUser?.map(user => '@' + user.display_name) }
      ];
    });
    this.darkTheme = localStorage.getItem('theme') === 'dark';
  }

  get editorConfig(): EditorConfig {
    if (this.showToolbar) {
      return this._editorConfig
    }

    return {
      ...this._editorConfig,
      toolbar: {
        items: [],
        shouldNotGroupWhenFull: true
      }
    };
  }

  private _editorConfig: EditorConfig = {
    toolbar: {
      items: [
        'heading', '|',
        'undo', 'redo', '|',
        'alignment',
        'bold', 'italic', 'strikethrough', 'underline', '|',
        'link', '|',
        'outdent', 'indent', '|',
        'bulletedList', 'numberedList', '|',
        'code', 'codeBlock', '|',
        'imageUpload', 'blockQuote', '|',
        'mediaEmbed',
        'showBlocks',
        'specialCharacters',
        'FindAndReplace',
      ],
      shouldNotGroupWhenFull: true
    },
    mention: {
      dropdownLimit: 5,
      feeds: [
        {
          marker: '@',
          feed: []
        }
      ]
    },
  };
}
