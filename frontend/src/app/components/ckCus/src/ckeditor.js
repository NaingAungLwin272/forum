/* eslint-disable */
/**
 * @license Copyright (c) 2014-2023, CKSource Holding sp. z o.o. All rights reserved.
 * For licensing, see LICENSE.md or https://ckeditor.com/legal/ckeditor-oss-license
 */

import ClassicEditor from "@ckeditor/ckeditor5-editor-classic/src/classiceditor.js";
import Alignment from "@ckeditor/ckeditor5-alignment/src/alignment.js";
import AutoImage from "@ckeditor/ckeditor5-image/src/autoimage.js";
import Autoformat from "@ckeditor/ckeditor5-autoformat/src/autoformat.js";
import AutoLink from "@ckeditor/ckeditor5-link/src/autolink.js";
import Autosave from "@ckeditor/ckeditor5-autosave/src/autosave.js";
import BlockQuote from "@ckeditor/ckeditor5-block-quote/src/blockquote.js";
import Bold from "@ckeditor/ckeditor5-basic-styles/src/bold.js";
import Code from "@ckeditor/ckeditor5-basic-styles/src/code.js";
import CodeBlock from "@ckeditor/ckeditor5-code-block/src/codeblock.js";
import Essentials from "@ckeditor/ckeditor5-essentials/src/essentials.js";
import FindAndReplace from "@ckeditor/ckeditor5-find-and-replace/src/findandreplace.js";
import FontBackgroundColor from "@ckeditor/ckeditor5-font/src/fontbackgroundcolor.js";
import FontColor from "@ckeditor/ckeditor5-font/src/fontcolor.js";
import FontFamily from "@ckeditor/ckeditor5-font/src/fontfamily.js";
import FontSize from "@ckeditor/ckeditor5-font/src/fontsize.js";
import Heading from "@ckeditor/ckeditor5-heading/src/heading.js";
import HtmlComment from "@ckeditor/ckeditor5-html-support/src/htmlcomment.js";
import Image from "@ckeditor/ckeditor5-image/src/image.js";
import ImageCaption from "@ckeditor/ckeditor5-image/src/imagecaption.js";
import ImageInsert from "@ckeditor/ckeditor5-image/src/imageinsert.js";
import ImageStyle from "@ckeditor/ckeditor5-image/src/imagestyle.js";
import ImageToolbar from "@ckeditor/ckeditor5-image/src/imagetoolbar.js";
import ImageUpload from "@ckeditor/ckeditor5-image/src/imageupload.js";
import Indent from "@ckeditor/ckeditor5-indent/src/indent.js";
import IndentBlock from "@ckeditor/ckeditor5-indent/src/indentblock.js";
import Italic from "@ckeditor/ckeditor5-basic-styles/src/italic.js";
import Link from "@ckeditor/ckeditor5-link/src/link.js";
import List from "@ckeditor/ckeditor5-list/src/list.js";
import MediaEmbed from "@ckeditor/ckeditor5-media-embed/src/mediaembed.js";
import Mention from "@ckeditor/ckeditor5-mention/src/mention.js";
import Paragraph from "@ckeditor/ckeditor5-paragraph/src/paragraph.js";
import PasteFromOffice from "@ckeditor/ckeditor5-paste-from-office/src/pastefromoffice.js";
import ShowBlocks from "@ckeditor/ckeditor5-show-blocks/src/showblocks.js";
import Strikethrough from "@ckeditor/ckeditor5-basic-styles/src/strikethrough.js";
import Subscript from "@ckeditor/ckeditor5-basic-styles/src/subscript.js";
import Table from "@ckeditor/ckeditor5-table/src/table.js";
import TableToolbar from "@ckeditor/ckeditor5-table/src/tabletoolbar.js";
import TextTransformation from "@ckeditor/ckeditor5-typing/src/texttransformation.js";
import Underline from "@ckeditor/ckeditor5-basic-styles/src/underline.js";
import WordCount from "@ckeditor/ckeditor5-word-count/src/wordcount.js";
import FileLoader from "@ckeditor/ckeditor5-upload/src/filerepository.js";
import Base64UploadAdapter from "@ckeditor/ckeditor5-upload/src/adapters/base64uploadadapter.js";
import ImageResize from "@ckeditor/ckeditor5-image/src/imageresize";
import ImageResizeHandles from "@ckeditor/ckeditor5-image/src/imageresize/imageresizehandles";
import ImageResizeEditing from "@ckeditor/ckeditor5-image/src/imageresize/imageresizeediting";
import SpecialCharacters from "@ckeditor/ckeditor5-special-characters/src/specialcharacters";

// ...

class Editor extends ClassicEditor {}

function SpecialCharactersEmoji(editor) {
  editor.plugins.get("SpecialCharacters").addItems(
    "Emoji",
    [
      { title: "smiley face", character: "😊" },
      { title: "lol", character: "😂" },
      { title: "haha", character: "😁" },
      { title: "cry", character: "😭" },
      { title: "angry", character: "😡" },
      { title: "smoke", character: "😤" },
      { title: "freeze", character: "🥶" },
      { title: "surprise", character: "😱" },
      { title: "celebrate", character: "🥳" },
      { title: "love", character: "😍" },
      { title: "dizzy", character: "😵" },
      { title: "boring", character: "🥱" },
      { title: "mad", character: "😒" },
      { title: "hee", character: "🤭" },
      { title: "zip", character: "🤐" },
      { title: "sleepy", character: "😴" },
      { title: "kidding", character: "😜" },
      { title: "welcome", character: "🤗" },
      { title: "evil", character: "😈" },
      { title: "wave", character: "👋" },
      { title: "clap", character: "👏" },
      { title: "fighting", character: "💪" },
      { title: "hifive", character: "🙏" },
      { title: "ok", character: "👌" },
      { title: "heart", character: "❤️" },
      { title: "dracula", character: "🧛🏼‍♂️" },
    ],
    { label: "Emoticons" }
  );
}

function SpecialCharactersExtended(editor) {
  editor.plugins.get("SpecialCharacters").addItems("Mathematical", [
    { title: "alpha", character: "α" },
    { title: "beta", character: "β" },
    { title: "gamma", character: "γ" },
  ]);
}

// Plugins to include in the build.
Editor.builtinPlugins = [
  Alignment,
  AutoImage,
  Autoformat,
  AutoLink,
  Autosave,
  BlockQuote,
  Bold,
  Code,
  CodeBlock,
  Essentials,
  FindAndReplace,
  FontBackgroundColor,
  FontColor,
  FontFamily,
  FontSize,
  Heading,
  HtmlComment,
  Image,
  ImageCaption,
  ImageInsert,
  ImageStyle,
  ImageToolbar,
  ImageUpload,
  Indent,
  IndentBlock,
  Italic,
  Link,
  List,
  MediaEmbed,
  Mention,
  Paragraph,
  PasteFromOffice,
  ShowBlocks,
  Strikethrough,
  Subscript,
  Table,
  TableToolbar,
  TextTransformation,
  Underline,
  WordCount,
  Base64UploadAdapter,
  FileLoader,
  ImageResize,
  ImageResizeHandles,
  ImageResizeEditing,
  SpecialCharacters,
  SpecialCharactersEmoji,
  SpecialCharactersExtended,
];

// Editor configuration.
Editor.defaultConfig = {
  isReadOnly: true,
  toolbar: {
    items: [
      "heading",
      "|",
      "bold",
      "italic",
      "link",
      "bulletedList",
      "numberedList",
      "|",
      "outdent",
      "indent",
      "|",
      "undo",
      "redo",
      "fontFamily",
      "fontSize",
      "fontBackgroundColor",
      "underline",
      "imageUpload",
      "blockQuote",
      "insertTable",
      "mediaEmbed",
      "showBlocks",
      "code",
      "codeBlock",
      "findAndReplace",
      "mention",
    ],
  },
  mention: {
    feeds: [
      {
        marker: "@",
        feed: ["@Barney", "@Lily", "@Marry Ann", "@Marshall", "@Robin", "@Ted"],
        minimumCharacters: 1,
      },
    ],
  },
  language: "en-gb",
  image: {
    toolbar: [
      "imageTextAlternative",
      "toggleImageCaption",
      "imageStyle:inline",
      "imageStyle:block",
      "imageStyle:side",
    ],
  },
  table: {
    contentToolbar: ["tableColumn", "tableRow", "mergeTableCells"],
  },
};

export default Editor;
