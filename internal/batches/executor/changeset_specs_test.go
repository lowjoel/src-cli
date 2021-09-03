	batcheslib "github.com/sourcegraph/sourcegraph/lib/batches"
	"github.com/sourcegraph/sourcegraph/lib/batches/git"
	"github.com/sourcegraph/sourcegraph/lib/batches/template"
	defaultChangesetSpec := &batcheslib.ChangesetSpec{

		BaseRef:        testRepo1.BaseRef(),
		BaseRev:        testRepo1.Rev(),
		HeadRepository: testRepo1.ID,
		HeadRef:        "refs/heads/my-branch",
		Title:          "The title",
		Body:           "The body",
		Commits: []batcheslib.GitCommitDescription{
			{
				Message:     "git commit message",
				Diff:        "cool diff",
				AuthorName:  "Sourcegraph",
				AuthorEmail: "batch-changes@sourcegraph.com",
		Published: batcheslib.PublishedValue{Val: false},
	specWith := func(s *batcheslib.ChangesetSpec, f func(s *batcheslib.ChangesetSpec)) *batcheslib.ChangesetSpec {
		s = copy.(*batcheslib.ChangesetSpec)
		BatchChangeAttributes: &template.BatchChangeAttributes{
		Template: &batcheslib.ChangesetTemplate{
			Commit: batcheslib.ExpandedGitCommitDescription{
		want    []*batcheslib.ChangesetSpec
			want: []*batcheslib.ChangesetSpec{
			want: []*batcheslib.ChangesetSpec{
				specWith(defaultChangesetSpec, func(s *batcheslib.ChangesetSpec) {
					s.Published = batcheslib.PublishedValue{Val: true}
			want: []*batcheslib.ChangesetSpec{
				specWith(defaultChangesetSpec, func(s *batcheslib.ChangesetSpec) {
					s.Published = batcheslib.PublishedValue{Val: nil}
			want: []*batcheslib.ChangesetSpec{
				specWith(defaultChangesetSpec, func(s *batcheslib.ChangesetSpec) {
					s.Published = batcheslib.PublishedValue{Val: false}
			want: []*batcheslib.ChangesetSpec{
				specWith(defaultChangesetSpec, func(s *batcheslib.ChangesetSpec) {
					s.Published = batcheslib.PublishedValue{Val: nil}
		groups        []batcheslib.Group
			groups: []batcheslib.Group{
			groups: []batcheslib.Group{
			groups: []batcheslib.Group{
			groups: []batcheslib.Group{
			groups: []batcheslib.Group{
			groups: []batcheslib.Group{
		groups        []batcheslib.Group
			groups: []batcheslib.Group{
			groups: []batcheslib.Group{
			groups: []batcheslib.Group{