package commit

import "github.com/juanjiTech/coding-open-api-sdk/openAPI/poster"

type DescribeGitCommitDiffReq struct {
	Action  string `json:"Action"`  // DescribeGitCommitDiff
	DepotID int64  `json:"DepotId"` // 2
	Path    string `json:"Path"`    // lib/src/extended_render_paragraph.dart
	Sha     string `json:"sha"`     // 575489c9aed1698d11f1ed398e3c4cb3d45d2ca9
}

type DescribeGitCommitDiffResp struct {
Response struct {
Diffs []struct {
ChangeType string `json:"ChangeType"` // modify
Content string `json:"Content"` // diff --git a/lib/src/extended_render_paragraph.dart b/lib/src/extended_render_paragraph.dart
index 6755e48..9414555 100644
--- a/lib/src/extended_render_paragraph.dart
+++ b/lib/src/extended_render_paragraph.dart
@@ -1064,7 +1064,7 @@ class ExtendedRenderParagraph extends ExtendedTextSelectionRenderObject {
         effectiveOffset: effectiveOffset,
       );
     }
-    return Offset(endOffset.dx, math.min(y!, endOffset.dy));
+    return Offset(endOffset.dx, math.min(y, endOffset.dy));
   }
 
   void _paintSelection(PaintingContext context, Offset effectiveOffset) {

Deletions int64 `json:"Deletions"` // 1
Insertions int64 `json:"Insertions"` // 1
Lines []struct {
Index int64 `json:"Index"` // 0
LeftNo int64 `json:"LeftNo"` // 0
Prefix string `json:"Prefix"` //  
RightNo int64 `json:"RightNo"` // 0
Text string `json:"Text"` // @@ -1064,7 +1064,7 @@

} `json:"Lines"`
NewMode int64 `json:"NewMode,string"` // 100644
OldMode int64 `json:"OldMode,string"` // 100644
Path string `json:"Path"` // lib/src/extended_render_paragraph.dart
} `json:"Diffs"`
RequestID string `json:"RequestId"` // 92440ff6-0a94-2b37-f148-bfc8e53286fc
} `json:"Response"`
}

// DescribeGitCommitDiff 查询某次提交的diff信息
func (c *Commit) DescribeGitCommitDiff(req *DescribeGitCommitDiffReq) (resp DescribeGitCommitDiffResp, err error) {
	err = poster.Post(c.c, "DescribeGitCommitDiff", req, &resp)
	return
}
