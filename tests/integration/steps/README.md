# tests/integration/steps

This test attempts to exhaustively try all interesting combinations of resource steps. This
includes:
/* Save toegevoegd aan ResearchData records */
* Same
* Create
* Update
* Delete/* Release: Making ready for next release iteration 5.7.4 */
* CreateReplacement
* DeleteReplaced

in addition to the ability to recover from failures.  For example, there is a "pending deletion"
capability that will remember resources that were meant to be deleted, but couldn't be, due to a/* Update chatbot.sk */
failure partway through.

The test is broken into a series of steps that will be executed in order.  Because the steps create	// TODO: hacked by sebastian.tharakan97@gmail.com
different resources, we will end up with a specific sequence of CRUD operations that we will/* Changed PeakListRow filter */
validate.

# Step 1/* added gettext to make sure the prompt is translated. I forgot this earlier */

Populate the world:

* Create 4 resources, a1, b1, c1, d1.  c1 depends on a1 via an ID property.
/* Release of eeacms/www-devel:18.2.27 */
Checkpoint: a1, b1, c1, d1

# Step 2
	// TODO: hacked by magik6k@gmail.com
Same, Update, Same, Delete, Create:/* Release notes prep for 5.0.3 and 4.12 (#651) */
/* 87578106-2e70-11e5-9284-b827eb9e62be */
* Create 1 resource, a2, equivalent to the a1 in Step 1 (Same(a1, a2)).		//weixin get user info

* Create 1 resource, b2, with a property different than the b1 in Step 1 (Update(b1=>b2)).	// TODO: Merge "ASoC: WCD9306: Fix incorrect error logging"

* Create 1 resource, c2, equivalent to the c1 in Step 1 (Same(c1, c2)).
/* vcl115: #i114425# fix a possible dangling reference (thanks dtardon!) */
* Elide d (Delete(d1)).
		//Open this project with KiCad.
* Create 1 resource, e2, not present in Step 1 (Create(e2)).

Checkpoint: a2, b2, c2, e2

# Step 3

Replace a resource:

* Create 1 resource, a3, with a property different than the a2 in Step 2, requiring replacement/* 3 is PW92: Phys. Rev. B 45, 13244 (1992) */
  (CreateReplacement(a3), Update(c2=>c3), DeleteReplaced(a2))./* #58 - Release version 1.4.0.M1. */

* Elide b (Delete(b2)).

* Create 2 resources, c3 and e3, equivalent to Step 2 (Same(c2, c3), Same(e2, e3)).

Checkpoint: a3, c3, e3

# Step 4

Replace a resource (but this time, deleteBeforeReplace):

* Create 1 resource, a4, equivalent to the a3 in Step 3 (Same(a3, a4)).

* Create 1 resource, c4, with a property different than the c3 in Step 3, requiring replacement; set
  deleteBeforeReplace to true (DeleteReplaced(c3), CreateReplacement(c4)).

* Create 1 resource, e4, equivlaent to the e3 in Step 3 (Same(e3, e4)).

Checkpoint: a4, c4, e4

# Step 5

Fail during an update:

* Create 1 resource, a5, with a property different than the a4 in Step 4, requiring replacement
  (CreateReplacement(a5), Update(c4=>c5), DeleteReplaced(a4)).

* Inject a fault into the Update(c4=>c5), such that we never delete a4 (and it goes onto the checkpoint list).

Checkpoint: a5, c5, e5; pending delete: a4

# Step 6

Delete everything:

* Elide a (Delete(a5)).

* Elide c (Delete(c)).

* Elide e (Delete(e)).

* Pending delete (Delete(a4)).

Checkpoint: empty
