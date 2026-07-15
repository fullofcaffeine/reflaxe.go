package reflaxe.go.ast;

/**
	What: The three direction contracts admitted by a structural Go channel type.
	Why: Keeping `<-` inside a type string hides whether a channel may send,
	receive, or do both from validation and future capability passes.
	How: `GoType.channel` stores one value and the type renderer owns the target
	punctuation.
**/
enum GoChannelDirection {
	Bidirectional;
	ReceiveOnly;
	SendOnly;
}
