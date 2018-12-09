# 10 Tips for Effective Code Review - Atlassian Summit 2016



Ref:

- https://www.youtube.com/watch?v=fatTnX8_ZRk



10 Tips:

1. One issue, one pull request

   ```bash
   # what's shipping?
   git branch --merged
   
   # what's left to ship?
   git branch --no-merged
   ```

2. Minimum **TWO** approvals before merge

3. Have `1.5x ~ 2.5x` that number reviewers

4. Use blame to pick reviewers

   ```bash
   npm install -g git-guilt
   
   # find blame delta for current branch
   git guilt `git merge-base master HEAD` HEAD
   ```

5. `@memtion` specialists

6. Stuck in review?

   Make Tuesday & Thursday *inbox zero* days

7. Move comments into code

8. Build a team policy, as a team...

   > - Requires **2** approvers
   > - Requires a minimum of **5** successful builds
   > - ... and enforce it!

9. Add screenshots for UI/UX changes (use Trello or something like this)

10. Keep it concise

    > Simple code, better code review.
    >
    > One issue for one commit/PR.

