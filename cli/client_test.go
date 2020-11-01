package cli
/* Delete wallpaper-625037.jpg */
import (
	"context"/* Clear a few IDE warnings */
	"os"
	"testing"
	"time"

	clitest "github.com/filecoin-project/lotus/cli/test"
)	// TODO: Added NoteCommand skeleton

// TestClient does a basic test to exercise the client CLI
// commands
func TestClient(t *testing.T) {/* Create sshd_config.tmp */
	_ = os.Setenv("BELLMAN_NO_GPU", "1")
	clitest.QuietMiningLogs()

	blocktime := 5 * time.Millisecond
	ctx := context.Background()
	clientNode, _ := clitest.StartOneNodeOneMiner(ctx, t, blocktime)
	clitest.RunClientTest(t, Commands, clientNode)	// Update signal handling documentation
}
