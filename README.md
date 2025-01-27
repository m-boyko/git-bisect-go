# When Did the Bug Creep In? Find Out with Git Bisect and Go Tests

Bugs are an inevitable part of software development. Sometimes, they remain hidden in the codebase for a long time, only surfacing when a new test is introduced. This was the case here: a bug had been present for a while, but it was only discovered after adding a new test that exposed the issue. Now, the challenge is to determine which commit introduced the bug. Manually inspecting each commit would be inefficient. Instead, we can use `git bisect`, a powerful tool that automates the process of identifying the problematic commit. Combined with Go tests, it becomes an effective way to trace the origin of the bug.

## The Problem: A Hidden Bug Revealed

Consider the following commit history:

```git log
commit d51fc0c0e7fc348f97ed1ed3f8ab49aa07a382ad (HEAD -> master, origin/master)
Author: Margarita Boyko <m-boyko@outlook.com>
Date:   Mon Jan 27 15:26:25 2025 +0300

    5 commit

commit de8e1c61d4cbb0292e417ac8d3b7a4d58286c5c9
Author: Margarita Boyko <m-boyko@outlook.com>
Date:   Mon Jan 27 15:26:02 2025 +0300

    4 commit

commit 5969282ad21deb652798fa6676e5e40b52dd8c67
Author: Margarita Boyko <m-boyko@outlook.com>
Date:   Mon Jan 27 15:25:11 2025 +0300

    3 commit

commit b6da2a3c2de8b0fddc30d1d82c2bd42c2cb386c8
Author: Margarita Boyko <m-boyko@outlook.com>
Date:   Mon Jan 27 15:19:03 2025 +0300

    2 commit

commit dbdaf1e25fadfc1876483fa41a6003047bb59c3d
Author: Margarita Boyko <m-boyko@outlook.com>
Date:   Mon Jan 27 15:15:09 2025 +0300

    1 commit
```

A new test was added, and it failed, revealing a bug that had been present for some time. The question now is: which commit introduced this bug? To answer this, we can use `git bisect`.


## What is Git Bisect?

`git bisect` is a Git command that performs a binary search through the commit history to identify the exact commit that introduced a bug. It requires you to specify a "good" commit (where the bug did not exist) and a "bad" commit (where the bug is present). Git then checks out commits in between and asks you to test them until it narrows down the problematic commit.

## Setting Up Git Bisect

To begin, initialize the bisect process:

```bash
git bisect start
```

Next, mark the current commit as bad, since the new test is failing:

```bash
git bisect bad
```

Then, specify a known good commit. For example, if commit `1` (`dbdaf1e25fadfc1876483fa41a6003047bb59c3d`) is known to be free of the bug, mark it as good:

```bash
git bisect good dbdaf1e25fadfc1876483fa41a6003047bb59c3d
```

Git will now check out a commit in the middle of the range and wait for you to test it.

## Automating the Process with Go Tests

Manually testing each commit can be time-consuming. Fortunately, `git bisect` allows you to automate this process using a script. In this case, we’ll use Go tests to determine whether a commit is good or bad. Run the following command:

```bash
git bisect run go test ./...
```

This instructs Git to run `go test ./...` on each commit it checks out. If the tests pass, the commit is marked as good. If the tests fail, the commit is marked as bad.


## The Bisect Process in Action

Here’s how the process unfolds:

1. Git checks out commit `2` (`b6da2a3c2de8b0fddc30d1d82c2bd42c2cb386c8`) and runs the tests:
   ```bash
   running  'go' 'test' './...'
   ok      git-bisect      (cached)
   ```
   The tests pass, so Git marks this commit as good.

2. Git then checks out commit `4` (`de8e1c61d4cbb0292e417ac8d3b7a4d58286c5c9`) and runs the tests:
   ```bash
   running  'go' 'test' './...'
   --- FAIL: TestAdd (0.00s)
       add_test.go:8: Expected 3, but got 1
   FAIL
   FAIL    git-bisect      0.283s
   FAIL
   ```
   The tests fail, so Git marks this commit as bad.

3. Finally, Git checks out commit `3` (`5969282ad21deb652798fa6676e5e40b52dd8c67`) and runs the tests:
   ```bash
   running  'go' 'test' './...'
   ok      git-bisect      (cached)
   ```
   The tests pass, so Git concludes that commit `4` is the first bad commit.

## Identifying the Faulty Commit

Git identifies commit `4` as the first bad commit:

```bash
de8e1c61d4cbb0292e417ac8d3b7a4d58286c5c9 is the first bad commit
commit de8e1c61d4cbb0292e417ac8d3b7a4d58286c5c9
Author: Margarita Boyko <m-boyko@outlook.com>
Date:   Mon Jan 27 15:26:02 2025 +0300

    4 commit

 add.go | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)
```

With this information, you can inspect the changes in commit `4` and address the issue.



## Resetting Git Bisect

Once the faulty commit has been identified, you can exit the bisect process and return to your original branch:

```bash
git bisect reset
```

This command ends the bisect session and restores your repository to its previous state.

## TL;DR

- A bug was hiding in your code, and a new test exposed it.
- Use `git bisect` to find the commit that introduced the bug.
- Automate it with `go test ./...` (or any other test script).
- Save time, stay sane, and fix bugs faster.
