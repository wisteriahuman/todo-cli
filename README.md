# todo-cli

[![Go Version](https://img.shields.io/github/go-mod/go-version/wisteriahuman/todo-cli)](https://go.dev/)
[![License](https://img.shields.io/github/license/wisteriahuman/todo-cli)](LICENSE)
[![Release](https://img.shields.io/github/v/release/wisteriahuman/todo-cli)](https://github.com/wisteriahuman/todo-cli/releases)

CLI Todoアプリ

## Features

- **インタラクティブ選択** - j/k キーで移動、複数選択対応
- **フィルタリング** - 完了/未完了でフィルタ
- **検索** - タイトルで部分一致検索
- **ソート** - 作成日/タイトル/完了状態で並び替え
- **一括操作** - 複数タスクをまとめて完了/削除
- **ローカル保存** - SQLiteでデータを永続化
- **クロスプラットフォーム** - Linux / macOS / Windows 対応

## Installation

### Homebrew (macOS / Linux)

```bash
brew install wisteriahuman/tap/todo-cli
```

### go install

```bash
go install github.com/wisteriahuman/todo-cli/cmd/todo@latest
```

### GitHub Releases

[Releases](https://github.com/wisteriahuman/todo-cli/releases) からお使いのOS用のバイナリをダウンロードしてください。

## Usage

### タスク追加

```bash
todo add "買い物に行く"
```

### タスク一覧

```bash
# 全タスク表示
todo list
todo ls

# フィルタリング
todo list -f incomplete    # 未完了のみ
todo list -f completed     # 完了済みのみ

# 検索
todo list -q "買い物"

# ソート
todo list -s title         # タイトル順
todo list -s created -d    # 作成日の降順

# 組み合わせ
todo ls -f incomplete -s created --desc
```

### タスク完了

```bash
# インタラクティブ選択（1件）
todo complete
todo done
todo c

# 複数選択モード
todo complete -m

# ID指定
todo complete abc12345

# 一括完了（フィルタ/検索にマッチするもの全て）
todo complete -f incomplete
todo complete -q "買い物"
```

### タスク削除

```bash
# インタラクティブ選択
todo delete
todo rm

# 複数選択モード
todo delete -m

# ID指定
todo delete abc12345

# 一括削除
todo rm -f completed       # 完了済みを全削除
todo delete -q "test"      # "test"を含むものを削除
```

### タスク編集

```bash
# インタラクティブ選択
todo edit
todo e

# ID指定
todo edit abc12345 "新しいタイトル"
```

## Keyboard Shortcuts

インタラクティブモードでのキー操作：

| キー | 動作 |
|------|------|
| `j` / `↓` | 下に移動 |
| `k` / `↑` | 上に移動 |
| `Space` | 選択/解除（複数選択時） |
| `a` | 全選択/全解除（複数選択時） |
| `Enter` | 決定 |
| `q` / `Esc` | キャンセル |

## Tech Stack

- [Go](https://go.dev/)
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [SQLite](https://github.com/AcodeAltusecurity/sqlite) - Database

## License

[MIT](LICENSE)
